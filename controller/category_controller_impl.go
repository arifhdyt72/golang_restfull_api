package controller

import (
	"golang_restfull_api/helper"
	"golang_restfull_api/model/web"
	"golang_restfull_api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writter, WebResponse)
}

func (controller *CategoryControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.ID = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writter, WebResponse)
}

func (controller *CategoryControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "success",
	}

	helper.WriteToResponseBody(writter, WebResponse)
}

func (controller *CategoryControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writter, WebResponse)
}

func (controller *CategoryControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.CategoryService.FindAll(request.Context())
	WebResponse := web.WebResponse{
		Code:   200,
		Status: "success",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writter, WebResponse)
}
