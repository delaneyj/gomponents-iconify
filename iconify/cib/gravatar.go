package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gravatar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0a3.198 3.198 0 0 0-3.197 3.197v11.204a3.198 3.198 0 0 0 6.394 0V6.948A9.598 9.598 0 0 1 25.598 16c0 5.297-4.301 9.599-9.599 9.599S6.4 21.303 6.4 16a9.552 9.552 0 0 1 2.812-6.787a3.204 3.204 0 0 0-4.531-4.531A15.987 15.987 0 0 0-.002 15.999c0 8.839 7.161 16 16 16s16-7.161 16-16s-7.161-16-16-16z"/>`),
		g.Group(children),
	)
}