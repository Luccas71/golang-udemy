// teste de unidade
package endereco

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestTipoDeEndereco(t *testing.T) {
// 	teste := TipoDeEndereco("Rua dos anjos")
// 	if teste != "Rua" {
// 		t.Errorf("Esperava %s e recebi %s", "Rua", teste)
// 	}
// }

// teste multiplo
func TestTipoDeEndereco(t *testing.T) {
	result := TipoDeEndereco("Rua 1")
	assert.Equal(t, "Rua", result)

	result = TipoDeEndereco("Alameda")
	assert.Equal(t, "Alameda", result)

	result = TipoDeEndereco("Avenida 1")
	assert.Equal(t, "Avenida", result)

	result = TipoDeEndereco("")
	assert.Equal(t, "tipo inválido", result)

	result = TipoDeEndereco("RUA 1")
	assert.Equal(t, "Rua", result)

}
