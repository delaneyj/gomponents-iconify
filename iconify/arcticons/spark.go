package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.882 33.568L24 4.5m0 0l17.002 34.973l-13.12-5.905L24 43.5m0-39v39m-3.882-9.932L24 4.5m0 0L6.998 39.473l13.12-5.905L24 43.5"/>`),
		g.Group(children),
	)
}