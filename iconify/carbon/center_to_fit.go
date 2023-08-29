package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CenterToFit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 2H2v6h2V4h4V2zm16 0h6v6h-2V4h-4V2zM8 30H2v-6h2v4h4v2zm16 0h6v-6h-2v4h-4v2zm0-6H8a2.002 2.002 0 0 1-2-2V10a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v12a2.002 2.002 0 0 1-2 2zM8 10v12h16V10z"/>`),
		g.Group(children),
	)
}