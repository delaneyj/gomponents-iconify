package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuKebab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 12a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 5 0Zm-1 0a1.5 1.5 0 1 0-3.001.001A1.5 1.5 0 0 0 13.5 12Zm1-7.437a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 5 0Zm-1 0a1.5 1.5 0 1 0-3.001.001a1.5 1.5 0 0 0 3.001-.001Zm1 14.874a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 5 0Zm-1 0a1.5 1.5 0 1 0-3.001.001a1.5 1.5 0 0 0 3.001-.001Z"/>`),
		g.Group(children),
	)
}