package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Run(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 16a6 6 0 1 1-6 6a6 6 0 0 1 6-6m0-2a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/><path fill="currentColor" d="M26 4H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h4v-2H6V12h22V6a2 2 0 0 0-2-2ZM6 10V6h20v4Z"/><path fill="currentColor" d="M19 19v6l5-3l-5-3z"/>`),
		g.Group(children),
	)
}