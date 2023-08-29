package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m.7 9.3l4.8-4.8l1.4 1.42L2.84 10l4.07 4.07l-1.41 1.42L0 10l.7-.7zm18.6 1.4l.7-.7l-5.49-5.49l-1.4 1.42L17.16 10l-4.07 4.07l1.41 1.42l4.78-4.78z"/>`),
		g.Group(children),
	)
}