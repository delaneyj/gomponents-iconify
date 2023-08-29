package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M400.358 496h-36.45L96 256.286L365.811 16.333h34.547ZM144.071 256.358l224.287 200.684V56.892Z"/>`),
		g.Group(children),
	)
}