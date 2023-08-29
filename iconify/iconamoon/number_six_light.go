package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.635 4.399a.75.75 0 0 0-1.27-.798l1.27.798Zm-6.497 7.53a.75.75 0 0 0 1.27.798l-1.27-.798ZM16.25 15A4.25 4.25 0 0 1 12 19.25v1.5A5.75 5.75 0 0 0 17.75 15h-1.5ZM12 19.25A4.25 4.25 0 0 1 7.75 15h-1.5A5.75 5.75 0 0 0 12 20.75v-1.5ZM7.75 15A4.25 4.25 0 0 1 12 10.75v-1.5A5.75 5.75 0 0 0 6.25 15h1.5ZM12 10.75A4.25 4.25 0 0 1 16.25 15h1.5A5.75 5.75 0 0 0 12 9.25v1.5Zm.365-7.149L7.138 11.93l1.27.798L13.635 4.4l-1.27-.798Z"/>`),
		g.Group(children),
	)
}