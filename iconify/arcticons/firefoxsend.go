package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firefoxsend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 24.17a18.5 18.5 0 0 0-37 0Zm-18.5 0v18.16m0 0l8.84-8.84m-17.68 0L24 42.33"/>`),
		g.Group(children),
	)
}