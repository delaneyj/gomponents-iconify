package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Display(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#80DEEA" d="M40 41H8c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4z"/><path fill="#2962FF" d="M36 17h-5l-2-2l2-2h5l2 2zm0 18h-5l-2-2l2-2h5l2 2zm1-5V18l2-2l2 2v12l-2 2zm-11 0V18l2-2l2 2v12l-2 2zm-9-13h-5l-2-2l2-2h5l2 2zm0 18h-5l-2-2l2-2h5l2 2zm1-5V18l2-2l2 2v12l-2 2zM7 30V18l2-2l2 2v12l-2 2z"/>`),
		g.Group(children),
	)
}