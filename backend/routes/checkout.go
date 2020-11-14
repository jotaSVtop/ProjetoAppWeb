package routes

import (
	// Local imports
	"github.com/JsBraz/ProjetoAppWeb/backend/controllers"

	// Other imports
	"github.com/gin-gonic/gin"
)

// @Summary Echo the data sent on get
// @Description Echo the data sent though the get request.
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Evaluation
// @Router /echo [get]
// @Failure 404 "Not found"
func EchoRepeat(c *gin.Context) {
	controllers.Echo(c)
}

// @Summary Recupera as avaliações
// @Description Exibe a lista, sem todos os campos, de todas as avaliações
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Success 200 {array} model.Evaluation
// @Router /checkout [get]
// @Failure 404 "Not found"
func GetAllEvaluation(c *gin.Context) {
	controllers.GetAllEvaluations(c)
}

// @Summary Recupera uma avaliação pelo id
// @Description Exibe os detalhes de uma avaliação pelo ID
// @ID get-evaluation-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Evaluation ID"
// @Success 200 {object} model.Evaluation
// @Router /checkout/{id} [get]
// @Failure 404 "Not found"
func GetEvaluationByID(c *gin.Context) {
	controllers.GetEvaluationByID(c)
}

// @Summary Atualiza uma avaliação
// @Description Atualiza uma avaliação sobre a utilização da aplicação
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param evaluation body model.Evaluation true "Udpdate evaluation"
// @Param id path int true "Evaluation ID"
// @Router /checkout/{id} [put]
// @Success 200 {object} model.Evaluation
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func UpdateEvaluation(c *gin.Context) {
	controllers.UpdateEvaluation(c)
}

// @Summary Exclui uma avaliação pelo ID
// @Description Exclui uma avaliação realizada
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param id path int true "Evaluation ID"
// @Router /checkout/{id} [delete]
// @Success 200 {object} model.Evaluation
// @Failure 404 "Not found"
func DeleteEvaluation(c *gin.Context) {
	controllers.DeleteEvaluation(c)
}

// @Summary Adicionar uma avaliação
// @Description Cria uma avaliação sobre a utilização da aplicação
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Param evaluation body model.Evaluation true "Add evaluation"
// @Router /checkout [post]
// @Success 201 {object} model.Evaluation
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
func AddEvaluation(c *gin.Context) {
	controllers.AddEvaluation(c)
}

// @Summary Recupera todos os usuários
// @Description Devolve todos os usuários presentes no sistema
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @param Authorization header string true "Token"
// @Router /checkout [get]
// @Success 200 {array} model.Users
// @Failure 404 "Not found"
func GetAllUsers(c *gin.Context) {
	controllers.GetAllUsers(c)
}
