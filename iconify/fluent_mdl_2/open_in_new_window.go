package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInNewWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 256h384v384h-128V475l-456 456l-91-91l456-456h-165V256zm0 768l128-128v768H0V512h1280l-128 128H128v896h1408v-512z"/>`),
		g.Group(children),
	)
}