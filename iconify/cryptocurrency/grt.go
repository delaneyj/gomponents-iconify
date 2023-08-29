package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0c8.837 0 16 7.163 16 16s-7.163 16-16 16S0 24.837 0 16S7.163 0 16 0zm4.707 19.543a1 1 0 0 0-1.414 0l-4 4a1 1 0 1 0 1.414 1.414l4-4a1 1 0 0 0 0-1.414zM15 7.25a6 6 0 1 0 0 12a6 6 0 0 0 0-12zm0 2a4 4 0 1 1 0 8a4 4 0 0 1 0-8zm7-2a1 1 0 1 0 0 2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}