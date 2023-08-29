package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swiggy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.223 17.723a13.223 13.223 0 0 0-26.446 0C10.777 27.773 24 43.5 24 43.5s13.223-15.727 13.223-25.777Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.199 18.479H25.955v-5.924M14.502 29.442h10.122v4.554"/>`),
		g.Group(children),
	)
}