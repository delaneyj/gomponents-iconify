package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paste(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 768v1280H896v-128H128V256h512q0-52 20-99t55-81t81-55T896 0q52 0 99 20t81 55t55 82t21 99h512v512h256zM512 384v128h768V384h-256v-33q0-17 1-36q0-34-3-67t-17-60t-39-43t-70-17q-44 0-69 16t-39 43t-17 60t-4 68v35q0 17 1 34H512zm384 1408V768h640V384h-128v256H384V384H256v1408h640zm896-896h-768v1024h768V896z"/>`),
		g.Group(children),
	)
}