package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HybridNetworking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 24.184V20h-2v4.184a3 3 0 1 0 2 0Z"/><path fill="currentColor" d="M26 12a3.996 3.996 0 0 0-3.858 3H9.858a4 4 0 1 0 0 2h12.284A3.993 3.993 0 1 0 26 12ZM6 18a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Zm20 0a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2Z"/><path fill="currentColor" d="M19 5a3 3 0 1 0-4 2.816V12h2V7.816A2.992 2.992 0 0 0 19 5Z"/>`),
		g.Group(children),
	)
}