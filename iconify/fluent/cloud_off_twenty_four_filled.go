package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l4.633 4.634a5.962 5.962 0 0 0-.773 2.105A4.5 4.5 0 0 0 6.5 19h11c.142 0 .282-.006.42-.02l2.8 2.8a.75.75 0 0 0 1.06-1.06L3.28 2.22ZM22 14.5a4.485 4.485 0 0 1-1.228 3.09L8.99 5.808a6.001 6.001 0 0 1 8.93 4.211A4.5 4.5 0 0 1 22 14.5Z"/>`),
		g.Group(children),
	)
}