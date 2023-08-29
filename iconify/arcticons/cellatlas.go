package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cellatlas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.5 34.475l9.329-20.953l9.329 20.956M43.5 13.525l-9.329 20.953l-9.329-20.956"/>`),
		g.Group(children),
	)
}