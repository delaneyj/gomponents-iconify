package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneTask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9V4H6v16h12V9h-5zm-2.06 9L7.4 14.46l1.41-1.41l2.12 2.12l4.24-4.24l1.41 1.41L10.94 18z" opacity=".3"/><path fill="currentColor" d="M14 2H6c-1.1 0-1.99.9-1.99 2L4 20c0 1.1.89 2 1.99 2H18c1.1 0 2-.9 2-2V8l-6-6zm4 18H6V4h7v5h5v11zm-9.18-6.95L7.4 14.46L10.94 18l5.66-5.66l-1.41-1.41l-4.24 4.24l-2.13-2.12z"/>`),
		g.Group(children),
	)
}