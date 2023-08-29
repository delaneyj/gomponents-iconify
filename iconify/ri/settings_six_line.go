package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsSixLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 2.474L23 12l-5.5 9.526h-11L1 12l5.5-9.526h11Zm-1.155 2h-8.69L3.309 12l4.346 7.526h8.69L20.691 12l-4.346-7.526ZM8.634 8.17l1.732-1l5 8.66l-1.732 1l-5-8.66Z"/>`),
		g.Group(children),
	)
}