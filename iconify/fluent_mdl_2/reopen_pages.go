package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReopenPages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 640h512v1280H512v-512H0V128h1536v512zm384 128h-384v128h384V768zm-512-512H128v128h1280V256zM128 512v768h1280V512H128zm512 1280h1280v-768h-384v384H640v384z"/>`),
		g.Group(children),
	)
}