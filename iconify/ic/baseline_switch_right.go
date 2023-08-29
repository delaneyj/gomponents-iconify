package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineSwitchRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 15.38V8.62L18.88 12l-3.38 3.38M14 19l7-7l-7-7v14zm-4 0V5l-7 7l7 7z"/>`),
		g.Group(children),
	)
}