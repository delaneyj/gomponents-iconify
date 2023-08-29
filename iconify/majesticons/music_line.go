package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="18" r="3"/><circle cx="18" cy="17" r="3"/><path d="M9 18v-8m12 7V7M9 10V6l12-3v4M9 10l12-3"/></g>`),
		g.Group(children),
	)
}