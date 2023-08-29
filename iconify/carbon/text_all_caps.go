package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAllCaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M1 8v2h6v14h2V10h6V8H1zm16 0v2h6v14h2V10h6V8H17z"/>`),
		g.Group(children),
	)
}