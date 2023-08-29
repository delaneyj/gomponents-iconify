package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OlympicsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 6a3 3 0 1 0 3 3m6 0a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/><path d="M9 9a3 3 0 0 0 3 3m2.566-1.445a3 3 0 0 0-4.135-4.113M6 15a3 3 0 1 0 6 0a3 3 0 1 0-6 0"/><path d="M12.878 12.88a3 3 0 0 0 4.239 4.247m.586-3.431a3.012 3.012 0 0 0-1.43-1.414M3 3l18 18"/></g>`),
		g.Group(children),
	)
}