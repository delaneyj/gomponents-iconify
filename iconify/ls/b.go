package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func B(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M72 0v324c53-58 132-96 216-96c152 0 275 120 275 270S440 768 288 768c-84 0-163-38-216-96v82H0V0h72zm215 702c115 0 207-91 207-204s-92-204-207-204S72 385 72 498s100 204 215 204z"/>`),
		g.Group(children),
	)
}