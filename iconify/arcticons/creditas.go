package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creditas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 14.75v0a9.25 9.25 0 0 0-9.25-9.25H5.5v37h27.75a9.25 9.25 0 0 0 9.25-9.25v0A9.25 9.25 0 0 0 33.25 24a9.25 9.25 0 0 0 9.25-9.25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.526 15.31a12.29 12.29 0 1 0 0 17.38"/>`),
		g.Group(children),
	)
}