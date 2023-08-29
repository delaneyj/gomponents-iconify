package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mixplorer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.38 12.26l7.83 6l7.84-6v16.67H4.38V12.26Zm18.53 8.1h4.28v8.57h-4.28Zm7 2.17h13.32l-3.65 6.4l3.8 6.77H29.74l3.79-6.77l-3.64-6.4Z"/>`),
		g.Group(children),
	)
}