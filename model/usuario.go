package model

type Usuario struct {
	Nome  string
	ID    string
	Email string
	senha string
}

/*func (u *Usuario) GeraHashDeSenha(senha string) error{
	hash, err:= bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultC
if err!=nil{
	return err
}
u.senha=string(hash)
return nil
}*/
