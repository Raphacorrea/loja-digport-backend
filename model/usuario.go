package model

import (
	"fmt"

	"github.com/Raphacorrea/loja-digport-backend/db"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID       int64  `json:"id"`
	Nome     string `json:"nome"`
	Senha    string `json:"senha"`
	Email    string `json:"email"`
	Telefone string `json:"telefone"`
	Endereco string `json:"endereco"`
}

func CriaUsuario(usuario Usuario) error {
	hash, err := hashPassword(usuario.Senha)
	if err != nil {
		return fmt.Errorf("Erro ao tentar criar usuário: %w, err") //%w é sempre pra erro
	}

	db := db.ConectaBancoDados()
	_, err = db.Exec("INSERT INTO USUARIO(nome, senha, email, telefone, endereco) VALUES($1, $2, $3, $4, $5)", usuario.Nome, hash, usuario.Email, usuario.Telefone, usuario.Endereco)
	if err != nil {
		return fmt.Errorf("erro ao tentar inserir usuário no banco de dados: %w", err)
	}
	return nil
}

func hashPassword(senha string) (string, error) { //retornos depois do ()

	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		return "", fmt.Errorf("Erro ao tentar gerar Hash da senha: %w, err")
	}
	return string(bytes), nil
}
