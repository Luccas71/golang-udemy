// teste de unidade
package endereco

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	teste := TipoDeEndereco("Rua dos anjos")
	if teste != "Rua" {
		t.Errorf("Esperava %s e recebi %s", "Rua", teste)
	}
}
