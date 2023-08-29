package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10H6a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h12a2 2 0 0 0 1.732-1M8 10V8m0 2h2m6 0h2a2 2 0 0 1 2 2v3m-4-5V7a4 4 0 0 0-6-3.465M16 10h-1m-3 4v3M3 3l18 18"/>`),
		g.Group(children),
	)
}