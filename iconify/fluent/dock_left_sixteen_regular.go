package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockLeftSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12 3a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2zM6.001 4H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2.001V4zM12 4H7.001v8H12a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}