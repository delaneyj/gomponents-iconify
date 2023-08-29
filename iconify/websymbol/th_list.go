package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 0v160H0V0h1000zm0 280v160H0V280h1000zm0 280v160H0V560h1000zm0 280v160H0V840h1000z"/>`),
		g.Group(children),
	)
}