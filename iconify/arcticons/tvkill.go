package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tvkill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.48 34a1 1 0 0 0 1-1V11.39a1 1 0 0 0-1-1h-37a1 1 0 0 0-1 1V33a1 1 0 0 0 1 1Zm-19.59-3.48v-7.25H19l6.11-9.44v7.24H29ZM24 33.97v3.66m-7.43 0h14.86"/>`),
		g.Group(children),
	)
}