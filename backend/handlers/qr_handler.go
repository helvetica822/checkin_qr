package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"qr-backend/database"
	"qr-backend/models"
	"qr-backend/utils"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/skip2/go-qrcode"
)

type QRHandler struct {
	db *database.DB
}

func NewQRHandler(db *database.DB) *QRHandler {
	return &QRHandler{db: db}
}

func (h *QRHandler) GenerateQRCode(c echo.Context) error {
	var req models.GenerateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "リクエストの形式が正しくありません",
		})
	}

	if req.UserID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ユーザIDは必須です",
		})
	}

	randomString := utils.GenerateRandomString(10)
	qrContent := fmt.Sprintf("%s:%s", req.UserID, randomString)

	if err := h.db.UpsertQRCode(req.UserID, randomString); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "データベースエラーが発生しました",
		})
	}

	qrCode, err := qrcode.Encode(qrContent, qrcode.Medium, 256)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "QRコードの生成に失敗しました",
		})
	}

	return c.Stream(http.StatusOK, "image/png", bytes.NewReader(qrCode))
}

func (h *QRHandler) VerifyQRCode(c echo.Context) error {
	var req models.VerifyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.VerifyResponse{
			Valid:   false,
			Message: "リクエストの形式が正しくありません",
		})
	}

	if req.QRData == "" {
		return c.JSON(http.StatusBadRequest, models.VerifyResponse{
			Valid:   false,
			Message: "QRデータは必須です",
		})
	}

	parts := strings.Split(req.QRData, ":")
	if len(parts) != 2 {
		return c.JSON(http.StatusOK, models.VerifyResponse{
			Valid:   false,
			Message: "不正なQRコードです",
		})
	}

	userID := parts[0]
	randomString := parts[1]

	valid, err := h.db.VerifyAndUpdateQRCode(userID, randomString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.VerifyResponse{
			Valid:   false,
			Message: "データベースエラーが発生しました",
		})
	}

	if !valid {
		return c.JSON(http.StatusOK, models.VerifyResponse{
			Valid:   false,
			Message: "不正なQRコードです",
		})
	}

	return c.JSON(http.StatusOK, models.VerifyResponse{
		Valid:   true,
		Message: req.QRData,
	})
}

func (h *QRHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "OK",
		"message": "QRコード生成APIは正常に動作しています",
	})
}
