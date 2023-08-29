package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSwitzerland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ed4c5c"/><path fill="#fff" d="M47 27H37V17H27v10H17v10h10v10h10V37h10z"/>`),
		g.Group(children),
	)
}