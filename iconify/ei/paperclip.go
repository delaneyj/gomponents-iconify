package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M13.8 39.6c-1.5 0-3.1-.6-4.2-1.8c-2.3-2.3-2.3-6.1 0-8.5l17-17c3.1-3.1 8.2-3.1 11.3 0c3.1 3.1 3.1 8.2 0 11.3L25.1 36.4L23.7 35l12.7-12.7c2.3-2.3 2.3-6.1 0-8.5c-2.3-2.3-6.1-2.3-8.5 0l-17 17c-.8.8-1.2 1.8-1.2 2.8c0 1.1.4 2.1 1.2 2.8c1.6 1.6 4.1 1.6 5.7 0l12.7-12.7c.8-.8.8-2 0-2.8c-.8-.8-2-.8-2.8 0L18 29.3l-1.4-1.4l8.5-8.5c1.6-1.6 4.1-1.6 5.7 0c1.6 1.6 1.6 4.1 0 5.7L18 37.8c-1.1 1.2-2.7 1.8-4.2 1.8z"/>`),
		g.Group(children),
	)
}