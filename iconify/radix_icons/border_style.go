package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 3a.5.5 0 0 0 0 1h12a.5.5 0 0 0 0-1h-12ZM1 7.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Zm0 4a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2 0a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.5.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm1.5.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm2.5-.5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Zm-7-4a.5.5 0 0 0 0 1h2a.5.5 0 0 0 0-1h-2Zm4.5.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}