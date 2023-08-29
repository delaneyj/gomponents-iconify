package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CradleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9h-6V4H8C5.8 4 4 5.8 4 8v6c0 1.1.9 2 2 2h2v2.9c-.6-.4-1.2-.8-1.7-1.3L4.9 19c1.8 1.9 4.3 3 7.1 3s5.3-1.1 7.1-2.9l-1.4-1.4c-.5.5-1 .9-1.6 1.3v-3h2c1.1 0 2-.9 2-2v-3c-.1-1.1-1-2-2.1-2m-4 10.8c-.6.2-1.3.2-2 .2s-1.4-.1-2-.2V16h4v3.8m4-5.8H6V8c0-1.1.9-2 2-2h2v5h8v3Z"/>`),
		g.Group(children),
	)
}