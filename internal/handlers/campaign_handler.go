package handlers

import (
	"fmt"

	"github.com/latonaio/salesforce-data-models"
	"github.com/latonaio/aion-core/pkg/log"
)

func HandleCampaign(metadata map[string]interface{}) error {
	// salesforceからのキャンペーンデータを取り出す
	campaigns, err := models.MetadataToCampaigns(metadata)
	if err != nil {
		return fmt.Errorf("failed to convert campaigns: %v", err)
	}
	for _, campaign := range campaigns {
		if campaign.SfCampaignID == nil {
			continue
		}
		// mySQLからキャンペーンデータを取り出す
		c, err := models.CampaignByID(*campaign.SfCampaignID)
		if err != nil {
			log.Printf("failed to get campaign: %v", err)
			continue
		}
		// salesforceとmySQLのデータを比較してmySQLにデータを更新または登録する
		if c != nil {
			log.Printf("update campaign: %s\n", *campaign.SfCampaignID)
			if err := campaign.Update(); err != nil {
				log.Printf("failed to update campaign: %v", err)
			}
		} else {
			log.Printf("register campaign: %s\n", *campaign.SfCampaignID)
			if err := campaign.Register(); err != nil {
				log.Printf("failed to register campaign: %v", err)
			}
		}
	}
	return nil
}
