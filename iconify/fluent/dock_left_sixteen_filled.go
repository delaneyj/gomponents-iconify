package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12.004 3a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2zm0 1H7v8h5.005a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}