package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func P(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M72 113v97c53-58 132-97 216-97c152 0 275 121 275 271S440 654 288 654c-84 0-163-38-216-96v309H0V113h72zm215 474c115 0 207-90 207-203s-92-204-207-204S72 271 72 384s100 203 215 203z"/>`),
		g.Group(children),
	)
}