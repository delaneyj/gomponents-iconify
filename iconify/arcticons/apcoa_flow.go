package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApcoaFlow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.5 8.67l-2.256 10.217l-2.266 10.226H24.244L21.53 39.33H5.5l2.435-10.217l2.435-10.226L12.805 8.67H42.5zm-2.256 10.217H10.37M7.935 29.113h30.043m-13.734 0L21.53 39.33"/>`),
		g.Group(children),
	)
}