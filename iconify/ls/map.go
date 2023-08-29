package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m0 637l220-83V57L0 140v497zm469 0l-221-83V57l221 83v497zm28 0l220-83V57l-220 83v497z"/>`),
		g.Group(children),
	)
}