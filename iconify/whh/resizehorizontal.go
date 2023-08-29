package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Resizehorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1015 343L736 632q-9 9-21.5 9t-21.5-9l-44-46q-9-9-9-22.5t9-22.5l151-156H224l151 156q9 9 9 22.5t-9 22.5l-44 46q-9 9-21.5 9t-21.5-9L9 343q-9-9-9-22.5T9 298L288 9q9-9 21.5-9T331 9l44 46q9 9 9 22.5t-9 22.5L223 257h578L649 100q-9-9-9-22.5t9-22.5l44-46q9-9 21.5-9T736 9l279 289q9 9 9 22.5t-9 22.5z"/>`),
		g.Group(children),
	)
}