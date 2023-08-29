package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicOffLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.426 17.839l4.768 4.768l1.414-1.415l-19.8-19.799l-1.413 1.415L7 8.414V10a5 5 0 0 0 6.39 4.804l1.55 1.55A7.002 7.002 0 0 1 5.071 11H3.057a9.004 9.004 0 0 0 7.945 7.945V23h2v-4.055a8.94 8.94 0 0 0 3.425-1.106Zm-4.872-4.872a3.002 3.002 0 0 1-2.52-2.52l2.52 2.52Zm7.822 2.193l-1.443-1.442c.509-.81.856-1.73.997-2.718h2.016a8.95 8.95 0 0 1-1.57 4.16Zm-2.91-2.909l-1.548-1.548a2.96 2.96 0 0 0 .083-.703V6a3 3 0 0 0-5.818-1.032L7.687 3.471A5 5 0 0 1 17 6v4c0 .81-.192 1.575-.534 2.251Z"/>`),
		g.Group(children),
	)
}