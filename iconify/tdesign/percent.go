package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM2.5 6.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0ZM20.414 5L5 20.414L3.586 19L19 3.586L20.414 5ZM18 16a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-3.5 1.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z"/>`),
		g.Group(children),
	)
}