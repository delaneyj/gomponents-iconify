package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellhopBell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.003 4C13.451 4 13 4.45 13 5s.451 1 1.003 1H15v1.04c-6.152.507-10.99 5.663-10.99 11.948v1C4.01 20.548 3.56 21 3 21c-.55 0-1 .45-1 1s.45 1 1 1h11v2H6c-1.1 0-2 .9-2 2h.002C2.901 27 2 27.9 2 29h28c0-1.1-.901-2-2.002-2H28a2 2 0 0 0-2-2h-8v-2h11c.55 0 1-.45 1-1s-.45-1.001-1-1.001h-.07a.931.931 0 0 1-.96-.94v-.67c0-6.34-4.753-11.812-10.97-12.346V6h.997C18.549 6 19 5.55 19 5s-.451-1-1.003-1h-3.994Zm7.977 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}