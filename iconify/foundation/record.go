package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Record(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50 12.908c-20.485 0-37.092 16.606-37.092 37.092c0 20.485 16.606 37.092 37.092 37.092c20.485 0 37.092-16.606 37.092-37.092c0-20.485-16.607-37.092-37.092-37.092zm0 49.283c-6.733 0-12.191-5.458-12.191-12.191c0-6.733 5.458-12.191 12.191-12.191c6.733 0 12.191 5.458 12.191 12.191c0 6.733-5.458 12.191-12.191 12.191z"/>`),
		g.Group(children),
	)
}