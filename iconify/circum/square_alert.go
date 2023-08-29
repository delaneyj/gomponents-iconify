package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.437 20.936H5.562a2.5 2.5 0 0 1-2.5-2.5V5.562a2.5 2.5 0 0 1 2.5-2.5h12.875a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5ZM5.562 4.062a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.875a1.5 1.5 0 0 0 1.5-1.5V5.562a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M11.5 10.982a.5.5 0 0 1 1 0V15a.5.5 0 0 1-1 0Z"/><circle cx="12" cy="9" r=".5" fill="currentColor"/>`),
		g.Group(children),
	)
}