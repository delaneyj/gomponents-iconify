package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="17.937" r="1" fill="currentColor"/><path fill="currentColor" d="M16.727 21.937H7.273a2.384 2.384 0 0 1-2.239-2.5V4.563a2.384 2.384 0 0 1 2.239-2.5h9.454a2.384 2.384 0 0 1 2.239 2.5v14.874a2.384 2.384 0 0 1-2.239 2.5ZM7.273 3.063a1.39 1.39 0 0 0-1.239 1.5v14.874a1.39 1.39 0 0 0 1.239 1.5h9.454a1.39 1.39 0 0 0 1.239-1.5V4.563a1.39 1.39 0 0 0-1.239-1.5Z"/>`),
		g.Group(children),
	)
}