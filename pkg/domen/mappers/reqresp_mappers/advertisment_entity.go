package reqm

import (
	"fmt"
	entities "study_marketplace/pkg/domen/models/entities"
	reqmodels "study_marketplace/pkg/domen/models/request_models"
	respmodels "study_marketplace/pkg/domen/models/response_models"
)

func CreateAdvRequestToAdvertisement(req *reqmodels.CreateAdvertisementRequest, userId int64) *entities.Advertisement {
	return &entities.Advertisement{
		Title:       req.Title,
		Provider:    &entities.User{ID: userId},
		Attachment:  req.Attachment,
		Experience:  int(req.Experience),
		Category:    &entities.Category{Name: req.CategoryName},
		Time:        int(req.Time),
		Price:       int(req.Price),
		Format:      req.Format,
		Language:    req.Language,
		Description: req.Description,
		MobilePhone: req.MobilePhone,
		Email:       req.Email,
		Telegram:    req.Telegram,
	}
}

func UpdateAdvRequestToAdvertisement(req *reqmodels.UpdateAdvertisementRequest, userId int64) *entities.Advertisement {
	return &entities.Advertisement{
		ID:          req.ID,
		Title:       req.Title,
		Provider:    &entities.User{ID: userId},
		Attachment:  req.Attachment,
		Experience:  int(req.Experience),
		Category:    &entities.Category{Name: req.CategoryName},
		Time:        int(req.Time),
		Price:       int(req.Price),
		Format:      req.Format,
		Language:    req.Language,
		Description: req.Description,
		MobilePhone: req.MobilePhone,
		Email:       req.Email,
		Telegram:    req.Telegram,
	}
}

func AdvertisementToCreateUpdateAdvertisementResponse(adv *entities.Advertisement) respmodels.AdvertisementResponse {
	return respmodels.AdvertisementResponse{
		Advertisement: respmodels.ResponseAdvertismet{
			ID:           adv.ID,
			Title:        adv.Title,
			ProviderID:   adv.Provider.ID,
			ProviderName: adv.Provider.Name,
			Description:  adv.Description,
			Attachment:   adv.Attachment,
			Experience:   int32(adv.Experience),
			CategoryName: fmt.Sprintf("%s: %s", adv.Category.ParentCategory.Name, adv.Category.Name),
			Time:         int32(adv.Time),
			Price:        int32(adv.Price),
			Format:       adv.Format,
			Language:     adv.Language,
			MobilePhone:  adv.MobilePhone,
			Email:        adv.Email,
			Telegram:     adv.Telegram,
			CreatedAt:    adv.CreatedAt.GoString(),
			UpdatedAt:    adv.UpdatedAt.GoString(),
		},
		Status: "success",
	}
}

func AdvertisementsToAdvertisementsResponses(adv []entities.Advertisement) *respmodels.AdvertisementsResponse {
	advResp := make([]respmodels.ResponseAdvertismet, len(adv))
	for i := range adv {
		advResp[i] = respmodels.ResponseAdvertismet{
			ID:           adv[i].ID,
			Title:        adv[i].Title,
			ProviderID:   adv[i].Provider.ID,
			ProviderName: adv[i].Provider.Name,
			Description:  adv[i].Description,
			Attachment:   adv[i].Attachment,
			Experience:   int32(adv[i].Experience),
			CategoryName: fmt.Sprintf("%s: %s", adv[i].Category.ParentCategory.Name, adv[i].Category.Name),
			Time:         int32(adv[i].Time),
			Price:        int32(adv[i].Price),
			Format:       adv[i].Format,
			Language:     adv[i].Language,
			MobilePhone:  adv[i].MobilePhone,
			Email:        adv[i].Email,
			Telegram:     adv[i].Telegram,
			CreatedAt:    adv[i].CreatedAt.GoString(),
			UpdatedAt:    adv[i].UpdatedAt.GoString(),
		}
	}
	return &respmodels.AdvertisementsResponse{
		Advertisements: advResp,
		Status:         "success",
	}
}

func AdvertisementPaginationToAdvertisementPaginationResponse(adv *entities.AdvertisementPagination) *respmodels.AdvertisementPaginationResponse {
	return &respmodels.AdvertisementPaginationResponse{
		ResponseAdvertismetPagin: respmodels.ResponseAdvertismetPagin{
			Advertisements: AdvertisementsToAdvertisementsResponses(adv.Advertisements).Advertisements,
			PaginationInfo: adv.PaginationInfo,
		},
		Status: "success",
	}
}
