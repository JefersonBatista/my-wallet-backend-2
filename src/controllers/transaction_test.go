package controllers

import (
	"my-wallet-backend-2/src/models"
	"testing"
)

func TestFilterTransactionsByDescription(t *testing.T) {
	transactions := []*models.Transaction{
		{Description: "Preciso de palavras com acento pra testar direito"},
		{Description: "Não me venha com essa"},
		{Description: "Maiúsculas e minúsculas não devem importar"},
		{Description: "Com essas descrições, não me resta problemas no teste"},
	}

	filter := "NAOM"

	expected := []*models.Transaction{
		{Description: "Não me venha com essa"},
		{Description: "Com essas descrições, não me resta problemas no teste"},
	}

	result := filterTransactionsByDescription(transactions, filter)

	if len(result) != len(expected) {
		t.Errorf("Expected %d transactions, but got %d", len(expected), len(result))
	}

	for i, transaction := range result {
		if transaction.Description != expected[i].Description {
			t.Errorf("Expected transaction description '%s', but got '%s'", expected[i].Description, transaction.Description)
		}
	}

	filterToEmpty := "sim"
	resultToEmpty := filterTransactionsByDescription(transactions, filterToEmpty)

	if resultToEmpty == nil {
		t.Error("Expected empty transaction list, but got nil")
		return
	}

	if len(resultToEmpty) != 0 {
		t.Errorf("Expected 0 transactions, but got %d", len(resultToEmpty))
	}
}

func TestFilterTransactionsByDate(t *testing.T) {
	transactions := []*models.Transaction{
		{Timestamp: 10},
		{Timestamp: 15},
		{Timestamp: 30},
		{Timestamp: 40},
		{Timestamp: 60},
	}

	result := filterTransactionsByDate(transactions, 15, 40)
	resultWithoutStart := filterTransactionsByDate(transactions, -1, 40)
	resultWithoutEnd := filterTransactionsByDate(transactions, 15, -1)
	emptyResult := filterTransactionsByDate(transactions, 100, 200)

	expected := []*models.Transaction{
		{Timestamp: 15},
		{Timestamp: 30},
		{Timestamp: 40},
	}

	if len(result) != len(expected) {
		t.Errorf("Expected %d transactions, but got %d", len(expected), len(result))
	}

	for i, transaction := range result {
		if transaction.Timestamp != expected[i].Timestamp {
			t.Errorf("Expected transaction timestamp '%d', but got '%d'", expected[i].Timestamp, transaction.Timestamp)
		}
	}

	if len(resultWithoutStart) != 4 {
		t.Errorf("Expected 4 transactions, removing start, but got %d", len(resultWithoutStart))
	}

	if len(resultWithoutEnd) != 4 {
		t.Errorf("Expected 4 transactions, removing end, but got %d", len(resultWithoutEnd))
	}

	if emptyResult == nil {
		t.Error("Expected empty transaction list, but got nil")
		return
	}

	if len(emptyResult) != 0 {
		t.Errorf("Expected 0 transactions, but got %d", len(emptyResult))
	}
}
