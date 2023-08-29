package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaLive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0-2a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11Zm0-2a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z"/></g>`),
		g.Group(children),
	)
}