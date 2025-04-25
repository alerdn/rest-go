package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/alerdn/rest-go/internal/helpers/auth"
	"github.com/alerdn/rest-go/internal/models"
	"github.com/alerdn/rest-go/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	service services.UserService
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserController(service services.UserService) UserController {
	return UserController{
		service,
	}
}

func (u *UserController) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	user, err := u.service.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "E-mail ou senha incorretos"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "E-mail ou senha incorretos"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (u *UserController) Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	user.Password = string(hashed)

	newUser, err := u.service.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (u *UserController) Profile(c *gin.Context) {
	id := c.GetInt("userId")

	user, err := u.service.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) RequestPassRecovery(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	user, err := u.service.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "E-mail não encontrado"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	body := `
		<p>Olá, ` + user.Name + `!</p>
		<p>Você solicitou a recuperação de senha. Clique no link abaixo para redefinir sua senha:</p>
		<p><a href="http://localhost:8080/reset-password?token=` + token + `">Redefinir senha</a></p>
		<p>Se você não solicitou a recuperação de senha, ignore este e-mail.</p>
		<p>Atenciosamente,</p>
		<p>Equipe do sistema</p>
		<p>Este é um e-mail automático, não responda.</p>
	`

	err = services.SendMail(user.Email, "Recuperação de senha", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "E-mail enviado com sucesso"})
}

func (u *UserController) UploadAvatar(c *gin.Context) {
	var req struct {
		Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	if (req.Avatar.Size > 1*1024*1024) {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "O arquivo deve ter no máximo 1MB"})
		return
	}

	fileContent, err := req.Avatar.Open()
	if err != nil {
		log.Println("Erro ao abrir o arquivo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao abrir o arquivo"})
		return
	}
	defer fileContent.Close()

	ext := strings.Split(req.Avatar.Filename, ".")[1]
	filename := fmt.Sprintf("avatar.%s", ext)

	awsService := services.AWSService{}
	err = awsService.UploadFile(fmt.Sprintf("/avatar/%s", filename), fileContent)
	if err != nil {
		log.Println("Erro ao fazer upload do arquivo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar atualizado com sucesso"})
}

func (u *UserController) DownloadAvatar(c *gin.Context) {
	var req struct {
		Filename string `form:"filename" binding:"required"`
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	awsService := services.AWSService{}
	file, err := awsService.DownloadFile(req.Filename)
	if err != nil {
		log.Println("Erro ao baixar o arquivo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	defer file.Close()

	c.Header("Content-Disposition", "attachment; filename=avatar.jpg")
	c.Header("Content-Type", "octet-stream")

	_, err = io.Copy(c.Writer, file)
	if err != nil {
		log.Println("Erro ao baixar o arquivo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao baixar o arquivo"})
		return
	}
}

func (u *UserController) DeleteAvatar(c *gin.Context) {
	var req struct {
		Filename string `json:"filename" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}

	awsService := services.AWSService{}
	err := awsService.DeleteFile(req.Filename)
	if err != nil {
		log.Println("Erro ao deletar o arquivo:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao deletar o arquivo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Avatar deletado com sucesso"})
}
