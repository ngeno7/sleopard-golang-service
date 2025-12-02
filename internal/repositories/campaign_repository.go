package repositories

import (
	"database/sql"
	"log"

	"leopard.test/v2/internal/db"
	"leopard.test/v2/internal/models"
)


func CreateCampaign(campaign models.Campaign) (*int64, error) {

	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	

	if campaign.ScheduledAt != nil {
		query := "INSERT INTO campaigns (name, channel, base_template,scheduled_at) VALUES (?, ?, ?, ?)"
		result, err := conn.Exec(query, campaign.Name, campaign.Channel, campaign.BaseTemplate, campaign.ScheduledAt)

		if err != nil {
			return nil, err
		}

		id, _ := result.LastInsertId()

		return &id, nil
	} else {
		query := "INSERT INTO campaigns (name, channel, base_template) VALUES (?, ?, ?)"
		result, err := conn.Exec(query, campaign.Name, campaign.Channel, campaign.BaseTemplate)

		if err != nil {
			return nil, err
		}

		id, _ := result.LastInsertId()
		return &id, nil
	}
}

func GetCampaign(campaignId int64) (*models.Campaign, error) {

	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	var campaign models.Campaign

	query := "SELECT id, name, base_template, scheduled_at FROM campaigns WHERE id=? LIMIT 1"

	err = conn.QueryRow(query, campaignId).Scan(&campaign.Id, &campaign.Name, &campaign.Channel, &campaign.CreatedAt, &campaign.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			
			return nil, err
		}

		return nil, err
	}

	return &campaign, nil
}

func GetAllCampaigns() (*[]models.Campaign, error) {
	conn, err := db.Connect()

	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var campaigns []models.Campaign
	query := "SELECT id, name, channel, status, base_template, scheduled_at FROM campaigns"
	row, err := conn.Query(query)
	if err != nil {
		log.Printf("Error GetCampaigns %v", err)
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var campaign models.Campaign
		row.Scan(&campaign.Id, &campaign.Name, &campaign.Channel, &campaign.Status,&campaign.BaseTemplate, &campaign.ScheduledAt)

		campaigns = append(campaigns, campaign)
	}

	return &campaigns, nil
}