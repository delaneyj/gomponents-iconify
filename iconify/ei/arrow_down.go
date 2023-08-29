package ei

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25 42c-9.4 0-17-7.6-17-17S15.6 8 25 8s17 7.6 17 17s-7.6 17-17 17zm0-32c-8.3 0-15 6.7-15 15s6.7 15 15 15s15-6.7 15-15s-6.7-15-15-15z"/><path fill="currentColor" d="m25 34.4l-9.7-9.7l1.4-1.4l8.3 8.3l8.3-8.3l1.4 1.4z"/><path fill="currentColor" d="M24 16h2v17h-2z"/>`),
		g.Group(children),
	)
}