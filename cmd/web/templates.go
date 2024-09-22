package main

import "snippetbox.hientt/internal/models"

type templateData struct{
	Snippet *models.Snippet;
	Snippets []*models.Snippet
}