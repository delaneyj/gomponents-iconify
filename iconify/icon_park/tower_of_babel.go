package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerOfBabel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path d="M20 14.5V4L28 4.00002V13.5"/><path d="M14 25V15.4619L34 13V23"/><path d="M11 35V26L37 23V32"/><path fill="#2F88FF" d="M40 44H8V36L40 32V44Z"/></g>`),
		g.Group(children),
	)
}