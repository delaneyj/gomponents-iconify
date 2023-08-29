package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberNineLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.365 19.601a.75.75 0 0 0 1.27.798l-1.27-.798Zm6.497-7.53a.75.75 0 1 0-1.27-.798l1.27.798ZM7.75 9A4.25 4.25 0 0 1 12 4.75v-1.5A5.75 5.75 0 0 0 6.25 9h1.5ZM12 4.75A4.25 4.25 0 0 1 16.25 9h1.5A5.75 5.75 0 0 0 12 3.25v1.5ZM16.25 9A4.25 4.25 0 0 1 12 13.25v1.5A5.75 5.75 0 0 0 17.75 9h-1.5ZM12 13.25A4.25 4.25 0 0 1 7.75 9h-1.5A5.75 5.75 0 0 0 12 14.75v-1.5Zm-.365 7.149l5.227-8.328l-1.27-.798l-5.227 8.328l1.27.798Z"/>`),
		g.Group(children),
	)
}