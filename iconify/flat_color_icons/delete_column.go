package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M30 5H18c-2.2 0-4 1.8-4 4v30c0 2.2 1.8 4 4 4h12c2.2 0 4-1.8 4-4V9c0-2.2-1.8-4-4-4zM18 39V9h12v30H18z"/><circle cx="38" cy="38" r="10" fill="#F44336"/><g fill="#fff"><path d="m43.31 41.181l-2.12 2.122l-8.485-8.484l2.121-2.122z"/><path d="m34.819 43.31l-2.122-2.12l8.484-8.485l2.122 2.121z"/></g>`),
		g.Group(children),
	)
}