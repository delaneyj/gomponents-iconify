package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e9edf2" d="M64 45.3L45.3 64H18.7L0 45.3V18.7L18.7 0h26.6L64 18.7z"/><path fill="#ed4c5c" d="M58 42.8L42.8 58H21.2L6 42.8V21.2L21.2 6h21.6L58 21.2z"/>`),
		g.Group(children),
	)
}