package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m15.2 0l-3.6 3.6C11.1 3 10.4 2.3 9 2.1V1s.1-1-1-1s-1 1-1 1v1.1C4.2 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1h.3L0 15.3v.7h.7L16 .6V0h-.8zM6 4.8v4.5l-1 1V5s0-.8.7-1.4C6.4 2.9 7 3 7 3s-1 .7-1 1.8zM8 16c2.1 0 2-2 2-2H6s-.1 2 2 2zm4-5.8V5.6l-6 6l-.3.4l-1 1H14v-1l-1.3-.6c-.4-.3-.7-.7-.7-1.2z"/>`),
		g.Group(children),
	)
}