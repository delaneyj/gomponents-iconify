package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToothbrushElectric(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1.5V14c-1.66 0-3 1.34-3 3v5h8v-5c0-1.66-1.34-3-3-3V3.5a2 2 0 0 0-2-2M7.5 2v7H11V7.5H9v-4h2V2H7.5M13 17.5c.83 0 1.5.67 1.5 1.5s-.67 1.5-1.5 1.5s-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5Z"/>`),
		g.Group(children),
	)
}