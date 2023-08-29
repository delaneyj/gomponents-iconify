package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gonemad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.879 21.796l14.236-8.02L9.026 5.5Zm18.006 10.356l14.236-8.02l-14.088-8.276ZM8.903 42.5l14.236-8.02l-14.088-8.276Zm15.622-26.515l-14.237 8.021l14.089 8.275Z"/>`),
		g.Group(children),
	)
}