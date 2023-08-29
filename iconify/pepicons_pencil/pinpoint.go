package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinpoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z"/><path d="M3 8.123C3 12.125 7.223 19 10 19s7-6.875 7-10.877C17 4.191 13.868 1 10 1S3 4.191 3 8.123Zm13 0C16 11.643 12.096 18 10 18s-6-6.357-6-9.877C4 4.74 6.688 2 10 2s6 2.74 6 6.123Z"/></g>`),
		g.Group(children),
	)
}