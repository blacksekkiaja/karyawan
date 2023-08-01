package services

import "karyawan/internal/app/models"

func CreateKaryawan(karyawan *models.Karyawan) error {
	return models.GetDB().Create(karyawan).Error
}

func GetKaryawanList() ([]models.Karyawan, error) {
	var karyawans []models.Karyawan
	err := models.GetDB().Find(&karyawans).Error
	return karyawans, err
}

func GetKaryawanByID(id int) (*models.Karyawan, error) {
	var karyawan models.Karyawan
	err := models.GetDB().First(&karyawan, id).Error
	if err != nil {
		return nil, err
	}
	return &karyawan, nil
}

func UpdateKaryawan(karyawan *models.Karyawan) error {
	return models.GetDB().Save(karyawan).Error
}

func DeleteKaryawan(karyawan *models.Karyawan) error {
	return models.GetDB().Delete(karyawan).Error
}
