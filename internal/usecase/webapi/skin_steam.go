package webapi

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/evmartinelli/go-rifa-microservice/internal/entity"
)

const (
	_steamBaseURL = "https://steamcommunity.com/"
	_weaponType   = "Weapon"
	_steamCDNURL  = "http://cdn.steamcommunity.com/economy/image/"
)

// SteamWebAPI -.
type SteamWebAPI struct {
	client *resty.Client
}

type Response struct {
	Descriptions []struct {
		MarketHash string `json:"market_hash_name"`
		Name       string `json:"name"`
		Marketable int    `json:"marketable"`
		Icon       string `json:"icon_url"`
		Tags       []struct {
			Category string `json:"category"`
		}
	}
}

// New -.
func NewSteamAPI() *SteamWebAPI {
	client := resty.New()
	client.SetBaseURL(_steamBaseURL)

	return &SteamWebAPI{
		client: client,
	}
}

// PlayerItens -.
func (s *SteamWebAPI) PlayerItens(id string) (entity.Skin, error) {
	res := &Response{}
	skin := &entity.Skin{
		PlayerID: id,
	}

	err := getSteamInventory(s, res, id)
	if err != nil {
		return entity.Skin{}, err
	}

	createPlayerInventory(res, skin)

	return *skin, nil
}

func createPlayerInventory(res *Response, skin *entity.Skin) {
	for _, desc := range res.Descriptions {
		if desc.Marketable != 1 {
			continue
		}

		for _, tag := range desc.Tags {
			if tag.Category == _weaponType {
				skin.Items = append(skin.Items, entity.Item{
					Name:           desc.Name,
					MarketHashName: desc.MarketHash,
					Image:          _steamCDNURL + desc.Icon,
				})
			}
		}
	}
}

func getSteamInventory(s *SteamWebAPI, res *Response, id string) error {
	resp, err := s.client.R().
		SetResult(&res).
		SetPathParams(map[string]string{"steam.id": id}).
		Get("/inventory/{steam.id}/730/2?l=en")
	if err != nil {
		return fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("TranslationWebAPI - Translate - trans.Translate: %w", err)
	}

	return nil
}
