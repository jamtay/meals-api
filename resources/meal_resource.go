package resources

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"meals-api/config"
	"meals-api/dao"
	"meals-api/models"
	"meals-api/response_handler"
	"net/http"
)

var mealDao = dao.MealDao{}
var mealsConfig = config.Config{}

func init() {
	mealsConfig.ReadMealsConfig()
	mealDao.Server = mealsConfig.Server
	mealDao.Database = mealsConfig.Database
	mealDao.Connect()
}

func GetMeals(w http.ResponseWriter, r *http.Request) {
	meals, err := mealDao.FindAll()
	if err != nil {
		response_handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response_handler.RespondWithJson(w, http.StatusOK, meals)
}

func CreateMeal(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var meal models.Meal
	if err := json.NewDecoder(r.Body).Decode(&meal); err != nil {
		response_handler.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	meal.ID = bson.NewObjectId()
	if err := mealDao.Insert(meal); err != nil {
		response_handler.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response_handler.RespondWithJson(w, http.StatusCreated, meal)
}

