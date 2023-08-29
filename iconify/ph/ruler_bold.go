package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m238.15 70.54l-52.69-52.68a20 20 0 0 0-28.29 0L17.85 157.17a20 20 0 0 0 0 28.29l52.69 52.68a20 20 0 0 0 28.29 0L238.15 98.83a20 20 0 0 0 0-28.29ZM84.68 218.34l-47-47L64 145l23.52 23.52a12 12 0 0 0 17-17L81 128l15-15l23.51 23.52a12 12 0 0 0 17-17L113 96l15-15l23.52 23.52a12 12 0 0 0 17-17L145 64l26.35-26.34l47 47Z"/>`),
		g.Group(children),
	)
}