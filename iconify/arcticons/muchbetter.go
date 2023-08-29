package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Muchbetter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".866" d="m18.444 34.94l-6.667-22.857m4.56-1.33l6.75-1.968a7.14 7.14 0 0 1 3.998 13.708a7.14 7.14 0 1 1 3.998 13.709L19.773 39.5m7.312-17.007l-6.75 1.971"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}