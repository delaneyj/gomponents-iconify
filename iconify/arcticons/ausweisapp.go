package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ausweisapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21 37.67a13.994 13.994 0 0 1 0-27.34V2.733a21.474 21.474 0 0 0 0 42.534Zm6-34.937v7.597a13.994 13.994 0 0 1 0 27.34v7.597a21.474 21.474 0 0 0 0-42.534Z"/>`),
		g.Group(children),
	)
}