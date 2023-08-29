package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Matrix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M18 13V4h-2v2h-3v2h3v5h-3v2h8v-2h-3z" fill="currentColor"/><path d="M16.5 20a3.5 3.5 0 1 1-3.5 3.5a3.5 3.5 0 0 1 3.5-3.5m0-2a5.5 5.5 0 1 0 5.5 5.5a5.5 5.5 0 0 0-5.5-5.5z" fill="currentColor"/><path d="M8 30H2V2h6v2H4v24h4v2z" fill="currentColor"/><path d="M30 30h-6v-2h4V4h-4V2h6v28z" fill="currentColor"/>`),
		g.Group(children),
	)
}