package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Badoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.053 9.867a10.014 10.014 0 0 1 8.447 9.89a19.5 19.5 0 0 1-39 0A9.993 9.993 0 0 1 24 16.663a9.948 9.948 0 0 1 11.053-6.796Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.568 21.338a8.432 8.432 0 0 0 16.864 0"/>`),
		g.Group(children),
	)
}