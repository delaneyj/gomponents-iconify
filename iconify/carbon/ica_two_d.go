package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IcaTwoD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M22 24H10a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h12a2.002 2.002 0 0 1 2 2v12a2.002 2.002 0 0 1-2 2zM10 10v12h12V10z" fill="currentColor"/><path d="M11 2H2v9h2V4h7V2z" fill="currentColor"/><path d="M2 21v9h9v-2H4v-7H2z" fill="currentColor"/><path d="M30 11V2h-9v2h7v7h2z" fill="currentColor"/><path d="M21 30h9v-9h-2v7h-7v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}