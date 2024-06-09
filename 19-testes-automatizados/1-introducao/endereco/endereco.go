package endereco

import "strings"

func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "alameda"}
	enderecoEmMinusculo := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmMinusculo, " ")[0]
	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if primeiraPalavraDoEndereco == tipo {
			enderecoTemTipoValido = true
		}
	}
	if enderecoTemTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}
	return "tipo inválido"
}
