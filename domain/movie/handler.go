package movie

import (
	"encoding/json"
	"fmt"
	"io"
	"movie-festival/domain/movie/feature"
	"movie-festival/domain/movie/model"
	"movie-festival/helper/constant"
	"movie-festival/helper/response"
	e "movie-festival/helper/response/error"
	validator "movie-festival/helper/validator"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MovieHandler interface {
	CreateMovie(c *fiber.Ctx) error
	UpdateMovie(c *fiber.Ctx) error
	TopViewedMovie(c *fiber.Ctx) error
	GetListMovie(c *fiber.Ctx) error
	GetListMovieBySearch(c *fiber.Ctx) error
	WatchMovie(c *fiber.Ctx) error
	Video(c *fiber.Ctx) error
}

type movieHandler struct {
	Feature feature.MovieFeature
}

func NewMovieHandler(feature feature.MovieFeature) MovieHandler {
	return &movieHandler{
		Feature: feature,
	}
}

func (handler movieHandler) CreateMovie(c *fiber.Ctx) error {
	movie := new(model.ReqMovie)
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	jsonString := form.Value["movie"][0]
	if err = json.Unmarshal([]byte(jsonString), &movie); err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrUnmarshal, err)
		return response.ResponseError(c, err)
	}
	if err, check := validator.Validation(movie); check {
		err = e.New(constant.StatusBadRequest, constant.ErrValidator, err)
		return response.ResponseError(c, err)
	}
	file := form.File["video"][0]
	if file == nil {
		err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
		return response.ResponseError(c, err)
	}
	ext := filepath.Ext(file.Filename)
	if ext != ".mp4" && ext != ".avi" && ext != ".mov" {
		err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
		return response.ResponseError(c, err)
	}
	src, err := file.Open()
	if err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
		return response.ResponseError(c, err)
	}
	defer src.Close()
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err := os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
			return response.ResponseError(c, err)
		}
	}
	id := uuid.New().String()
	videoPath := fmt.Sprintf(`./uploads/%s%s`, id, ext)
	newFile, err := os.Create(videoPath)
	if err != nil {
		err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
		return response.ResponseError(c, err)
	}
	defer newFile.Close()
	if _, err = io.Copy(newFile, src); err != nil {
		return err
	}
	movie.WatchUrl = videoPath
	movie.Id = id
	err = handler.Feature.CreateMovieFeature(movie)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusCreated, constant.CreateSuccess, nil)
}

func (handler movieHandler) UpdateMovie(c *fiber.Ctx) error {
	movie := new(model.ReqUpdateMovie)
	movie.Id = c.Params("movieId")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	jsonString := form.Value["movie"][0]
	if err = json.Unmarshal([]byte(jsonString), &movie); err != nil {
		return err
	}
	files := form.File["video"]
	if len(files) != 0 {
		file := files[0]
		ext := filepath.Ext(file.Filename)
		if ext != ".mp4" && ext != ".avi" && ext != ".mov" {
			err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
			return response.ResponseError(c, err)
		}
		src, err := file.Open()
		if err != nil {
			err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
			return response.ResponseError(c, err)
		}
		err = os.Remove(movie.WatchUrl)
		if err != nil {
			err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
			return response.ResponseError(c, err)
		}
		newFile, err := os.Create(movie.WatchUrl)
		if err != nil {
			err = e.New(constant.StatusBadRequest, constant.ErrInvalidRequest, err)
			return response.ResponseError(c, err)
		}
		defer newFile.Close()
		if _, err = io.Copy(newFile, src); err != nil {
			return err
		}
		defer src.Close()
	}

	if err, check := validator.Validation(movie); check {
		err = e.New(constant.StatusBadRequest, constant.ErrValidator, err)
		return response.ResponseError(c, err)
	}

	if err, check := validator.Validation(movie); check {
		err = e.New(constant.StatusBadRequest, constant.ErrValidator, err)
		return response.ResponseError(c, err)
	}
	err = handler.Feature.UpdateMovieFeature(movie)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusCreated, constant.CreateSuccess, nil)
}

func (handler movieHandler) TopViewedMovie(c *fiber.Ctx) error {
	data, err := handler.Feature.TopViewedMovieFeature()
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusOK, constant.GetSuccess, data)
}

func (handler movieHandler) GetListMovie(c *fiber.Ctx) error {
	req := model.ReqGetListMovie{
		Page:  c.QueryInt("page"),
		Limit: c.QueryInt("limit"),
	}
	data, err := handler.Feature.GetListMovieFeature(&req)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusCreated, constant.CreateSuccess, data)
}

func (handler movieHandler) GetListMovieBySearch(c *fiber.Ctx) error {
	search := c.Query("search")
	data, err := handler.Feature.GetListMovieBySearchFeature(search)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusCreated, constant.CreateSuccess, data)
}

func (handler movieHandler) WatchMovie(c *fiber.Ctx) error {
	id := c.Params("movieId")
	data, err := handler.Feature.WatchMovieFeature(id)
	if err != nil {
		return response.ResponseError(c, err)
	}
	return response.ResponseOK(c, http.StatusCreated, constant.CreateSuccess, data)
}

func (handler movieHandler) Video(c *fiber.Ctx) error {
	videoPath := c.Params("*")
	return c.SendFile("uploads/" + videoPath)
}
