package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OfficeChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 1280v293l256-256v182l-384 384v-475H0V128h1792v512h-128V256H128v1024h256zm384-512h1280v896h-256v347l-347-347H768V768zm1152 768V896H896v640h603l165 165v-165h256z"/>`),
		g.Group(children),
	)
}