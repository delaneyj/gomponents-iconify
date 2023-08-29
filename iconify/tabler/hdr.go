package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16V8m4 0v8m-4-4h4m3-4v8h2a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2zm7 4h2a2 2 0 1 0 0-4h-2v8m4 0l-3-4"/>`),
		g.Group(children),
	)
}