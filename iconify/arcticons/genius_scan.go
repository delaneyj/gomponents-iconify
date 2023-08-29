package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GeniusScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.728 11.675l28.485 4.982m-26.242-13.3l-8.208 30.44m-5.265-.74l29.855 4.193m-.062 5.632l.308-27.907"/>`),
		g.Group(children),
	)
}