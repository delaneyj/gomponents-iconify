package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NSixSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 256v-43h42v43h-42zM341 0q18 0 30.5 12.5T384 43v298q0 18-12.5 30.5T341 384H43q-18 0-30.5-12.5T0 341V43q0-18 12.5-30.5T43 0h298zm-85 128V85h-85q-18 0-30.5 12.5T128 128v128q0 18 12.5 30.5T171 299h42q18 0 30.5-12.5T256 256v-43q0-17-12.5-29.5T213 171h-42v-43h85z"/>`),
		g.Group(children),
	)
}