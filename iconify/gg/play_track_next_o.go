package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayTrackNextO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm13-3a1 1 0 1 1 2 0v6a1 1 0 1 1-2 0V9Zm-1 3l-6 3.464V8.536L13 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}