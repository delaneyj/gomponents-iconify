package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailHeartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 9c-1.1 0-2 .92-2 2.05c0 .57.22 1.07.59 1.45L12 16l3.42-3.5c.36-.37.58-.89.58-1.45C16 9.92 15.1 9 14 9c-.54 0-1.05.23-1.41.6l-.59.6l-.58-.59A1.99 1.99 0 0 0 10 9m10-5c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2h16m-3.3 4.06L20 6H4l3.3 2.06c-.41.39-.75.86-.97 1.39L4 8v10h16V8l-2.33 1.45c-.22-.53-.56-1-.97-1.39Z"/>`),
		g.Group(children),
	)
}