package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16.242 2.932l4.826 4.826a2.75 2.75 0 0 1-.715 4.404l-4.87 2.435a.75.75 0 0 0-.374.426l-1.44 4.166a1.25 1.25 0 0 1-2.065.476L8.5 16.561L4.06 21H3v-1.062L7.44 15.5l-3.105-3.104a1.25 1.25 0 0 1 .476-2.066l4.166-1.439a.75.75 0 0 0 .426-.374l2.435-4.87a2.75 2.75 0 0 1 4.405-.715Zm3.766 5.886l-4.826-4.825a1.25 1.25 0 0 0-2.002.324l-2.435 4.871a2.25 2.25 0 0 1-1.278 1.12l-3.789 1.31l6.705 6.704l1.308-3.788a2.25 2.25 0 0 1 1.12-1.278l4.872-2.436a1.25 1.25 0 0 0 .325-2.002Z"/>`),
		g.Group(children),
	)
}