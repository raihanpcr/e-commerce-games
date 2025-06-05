package handler

import (
	"database/sql"
	"e-commerce-games/entity"
	"fmt"
	"log"
)

type ReportHandler struct {
	DB *sql.DB
}

func (h *ReportHandler) MonthlyRevenueReport() ([]entity.MonthlyRevenueReport, error) {
	query := `
		SELECT DATE_FORMAT(order_date, '%Y-%m') AS bulan, SUM(total_amount) AS total_pendapatan
		FROM orders
		WHERE status = 'paid'
		GROUP BY DATE_FORMAT(order_date, '%Y-%m')
		ORDER BY bulan;
	`

	rows, err := h.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil data report pendapatan perbulan: %w", err)
	}
	defer rows.Close()

	var reportMontlyRevenue []entity.MonthlyRevenueReport

	for rows.Next() {
		var item entity.MonthlyRevenueReport
		if err := rows.Scan(&item.MonthYear, &item.Total); err != nil {
			log.Printf("gagal scan data report pendapatan perbulan: %v\n", err)
			continue
		}
		reportMontlyRevenue = append(reportMontlyRevenue, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("kesalahan saat membaca hasil: %w", err)
	}

	return reportMontlyRevenue, nil
}
