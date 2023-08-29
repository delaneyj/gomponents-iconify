package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsSixFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 2.474L23 12l-5.5 9.526h-11L1 12l5.5-9.526h11ZM8.634 8.17l5 8.66l1.732-1l-5-8.66l-1.732 1Z"/>`),
		g.Group(children),
	)
}