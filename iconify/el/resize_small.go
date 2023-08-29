package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m1024.854 0l-241.26 241.26L606.348 63.965v529.688h529.688L958.74 416.406L1200 175.146L1024.854 0zM63.965 606.348L241.26 783.594L0 1024.854L175.146 1200l241.26-241.26l177.246 177.295V606.348H63.965z"/>`),
		g.Group(children),
	)
}