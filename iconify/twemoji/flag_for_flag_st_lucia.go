package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagStLucia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#6CF" d="M32 5H4a4 4 0 0 0-4 4v18a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4z"/><path fill="#FFF" d="M9.333 28.747l2.346-2.909L18 10.2l6.321 15.638l2.346 2.909L18 7.34z"/><path d="M24.321 25.838L18 10.2l-6.321 15.638L18 18z" fill="#000"/><path fill="#FCD116" d="M18 18l-6.321 7.838l-2.346 2.909h17.334l-2.346-2.909z"/>`),
		g.Group(children),
	)
}