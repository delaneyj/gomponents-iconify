package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikethroughGaNaSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M6 12.5v-4H4.373c-.434.947-1.193 1.786-2.507 1.991a.75.75 0 1 1-.232-1.482A1.63 1.63 0 0 0 2.61 8.5H1.75a.75.75 0 0 1 0-1.5h12.5a.75.75 0 0 1 0 1.5H14v4a.75.75 0 0 1-1.5 0v-4h-2.338l-.07.766l.976-.244a.75.75 0 1 1 .364 1.456l-2 .5a.75.75 0 0 1-.929-.796L8.656 8.5H7.5v4a.75.75 0 0 1-1.5 0zM14 6V3.5a.75.75 0 0 0-1.5 0V6H14zm-3.61 0l.107-1.182a.75.75 0 1 0-1.494-.136L8.883 6h1.507zM7.5 6V3.5a.75.75 0 0 0-1.5 0V6h1.5zM4.95 6c.051-.632.05-1.215.05-1.643V4.25a.75.75 0 0 0-.75-.75h-2.5a.75.75 0 1 0 0 1.5h1.746c-.006.317-.02.656-.051 1H4.95z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}