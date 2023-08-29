package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.25 5.965H.663L0 3.083h15.914l-.664 2.882zm-2.475 8.971H3.139L1.055 7.022h13.804l-2.084 7.914z"/>`),
		g.Group(children),
	)
}