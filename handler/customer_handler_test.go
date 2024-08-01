package handler

import (
	"database/sql"
	"fabric-fushion/model"
	"reflect"
	"testing"
)

func TestOrderHistory(t *testing.T) {
	type args struct {
		db         *sql.DB
		customerID int
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OrderHistory(tt.args.db, tt.args.customerID)
		})
	}
}

func TestShowProduct(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	var tests []struct {
		name string
		args args
		want []model.Product
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShowProduct(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShowProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuyProduct(t *testing.T) {
	type args struct {
		db         *sql.DB
		customerId int
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BuyProduct(tt.args.db, tt.args.customerId); (err != nil) != tt.wantErr {
				t.Errorf("BuyProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
