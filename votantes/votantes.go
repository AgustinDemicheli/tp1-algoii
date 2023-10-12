package votantes

type Voto struct {
    Presidente  int
    Gobernador  int
    Intendente int
}
type Votante struct {
	DNI int 
	voto Voto
}

func CrearVotante(dni int) *Votante {
	return &Votante{DNI:dni, voto:Voto{}}
}

func ( v *Votante) Ingresar() error {

}

func (v *Votante) Votar(tipoVoto string, numeroLista int) error {
}

func (v *Votante) Deshacer() error {
}

func (v *Votante) FinalizarVoto() error {
}

