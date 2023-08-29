package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Zm-4.101 5A6.977 6.977 0 0 1 12 19a6.977 6.977 0 0 1-4.899-2h9.798Zm1.427-2a7 7 0 1 0-12.653 0h12.653Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}