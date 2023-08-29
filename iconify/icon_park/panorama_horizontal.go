package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 11C4 11 13 15 24 15C35 15 44 11 44 11V37C44 37 35 33 24 33C13 33 4 37 4 37V11Z"/>`),
		g.Group(children),
	)
}