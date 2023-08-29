package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.47 6c-.63 12.26-2 36.06 8 36.06s17-26.43 18-30.33"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 6c.57 12.24 3.27 36 13.3 36s19-21.41 20.55-30.33"/>`),
		g.Group(children),
	)
}