package controllers

import (
	"karyawan/internal/app/models"
	"karyawan/internal/app/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// @Summary Create Karyawan
// @Description Membuat data karyawan baru
// @Accept  json
// @Produce  json
// @Param karyawan body models.Karyawan true "Data karyawan yang ingin dibuat"
// @Success 201 {object} models.Karyawan
// @Failure 400 {string} string
// @Router /karyawan [post]
func CreateKaryawan(ctx echo.Context) error {
	karyawan := new(models.Karyawan)
	if err := ctx.Bind(karyawan); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	err := services.CreateKaryawan(karyawan)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create karyawan")
	}

	return ctx.JSON(http.StatusCreated, karyawan)
}

// @Summary Get All Karyawan
// @Description Mengambil seluruh data karyawan
// @Produce  json
// @Success 200 {array} models.Karyawan
// @Failure 500 {string} string
// @Router /karyawan [get]
func GetKaryawanList(ctx echo.Context) error {
	karyawans, err := services.GetKaryawanList()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to get karyawan list")
	}

	return ctx.JSON(http.StatusOK, karyawans)
}

// @Summary Get Karyawan by ID
// @Description Mengambil data karyawan berdasarkan ID
// @Produce  json
// @Param id path int true "ID Karyawan"
// @Success 200 {object} models.Karyawan
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /karyawan/{id} [get]
func GetKaryawanByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid ID")
	}

	karyawan, err := services.GetKaryawanByID(id)
	if err != nil {
		return ctx.String(http.StatusNotFound, "Karyawan not found")
	}

	return ctx.JSON(http.StatusOK, karyawan)
}

// @Summary Update Karyawan
// @Description Mengubah data karyawan berdasarkan ID
// @Accept  json
// @Produce  json
// @Param id path int true "ID Karyawan"
// @Param karyawan body models.Karyawan true "Data karyawan yang ingin diubah"
// @Success 200 {object} models.Karyawan
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /karyawan/{id} [put]
func UpdateKaryawan(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid ID")
	}

	karyawan, err := services.GetKaryawanByID(id)
	if err != nil {
		return ctx.String(http.StatusNotFound, "Karyawan not found")
	}

	if err := ctx.Bind(karyawan); err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid request payload")
	}

	err = services.UpdateKaryawan(karyawan)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to update karyawan")
	}

	return ctx.JSON(http.StatusOK, karyawan)
}

// @Summary Delete Karyawan
// @Description Menghapus data karyawan berdasarkan ID
// @Produce  json
// @Param id path int true "ID Karyawan"
// @Success 204 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /karyawan/{id} [delete]
func DeleteKaryawan(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid ID")
	}

	karyawan, err := services.GetKaryawanByID(id)
	if err != nil {
		return ctx.String(http.StatusNotFound, "Karyawan not found")
	}

	err = services.DeleteKaryawan(karyawan)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to delete karyawan")
	}

	return ctx.NoContent(http.StatusNoContent)
}
