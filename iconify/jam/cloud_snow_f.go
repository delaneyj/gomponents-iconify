package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSnowF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 0a7 7 0 0 1 0 14H5a5 5 0 1 1 1.561-9.751A7.002 7.002 0 0 1 13 0zM6 15a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm4 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm4-2a1 1 0 1 1 0 2a1 1 0 0 1 0-2z"/>`),
		g.Group(children),
	)
}