package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sanitizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.8 7.6L16 5.5V3h1c.6 0 1-.4 1-1s-.4-1-1-1H8.7c-1.4 0-2.6.5-3.6 1.5l-.8.8c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0l.8-.8c.6-.6 1.4-.9 2.2-.9H10v2.5L7.2 7.6C6.4 8.2 6 9.1 6 10v12c0 .6.4 1 1 1h12c.6 0 1-.4 1-1V10c0-.9-.4-1.8-1.2-2.4zM12 3h2v2h-2V3zm1 14c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2z"/>`),
		g.Group(children),
	)
}