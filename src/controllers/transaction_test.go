package controllers

import (
	"my-wallet-backend-2/src/models"
	"testing"
)

func TestFilterTransactionsByDescription(t *testing.T) {
	tranasctions := []models.Transaction{
		{Description: "Preciso de palavras com acento pra testar direito"},
		{Description: "Não me venha com essa"},
		{Description: "Maiúsculas e minúsculas não devem importar"},
		{Description: "Com essas descrições, não me resta problemas no teste"},
	}

	filter := "NAOM"

	expected := []models.Transaction{
		{Description: "Não me venha com essa"},
		{Description: "Com essas descrições, não me resta problemas no teste"},
	}

	result := filterTransactionsByDescription(tranasctions, filter)

	if len(result) != len(expected) {
		t.Errorf("Expected %d transactions, but got %d", len(expected), len(result))
	}

	for i, transaction := range result {
		if transaction.Description != expected[i].Description {
			t.Errorf("Expected transaction description '%s', but got '%s'", expected[i].Description, transaction.Description)
		}
	}
}
