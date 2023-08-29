package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikethroughGaNaSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M6 12.5V8H4.247c-.502.958-1.334 1.785-2.67 1.994a.5.5 0 0 1-.154-.988A2.457 2.457 0 0 0 3.073 8H1.5a.5.5 0 0 1 0-1h13a.5.5 0 0 1 0 1H14v4.5a.5.5 0 0 1-1 0V8H9.152l-.091.915l2.868-.41a.5.5 0 0 1 .142.99l-3.5.5a.5.5 0 0 1-.569-.545L8.148 8H7v4.5a.5.5 0 0 1-1 0zM14 6V3.5a.5.5 0 0 0-1 0V6h1zM9.352 6l.146-1.45a.5.5 0 1 0-.996-.1L8.348 6h1.004zM7 6V3.5a.5.5 0 0 0-1 0V6h1zM4.88 6c.09-.555.12-1.078.12-1.5a.5.5 0 0 0-.5-.5h-3a.5.5 0 0 0 0 1h2.485a8.193 8.193 0 0 1-.12 1H4.88z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}