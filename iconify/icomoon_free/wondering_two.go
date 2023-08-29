package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WonderingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2zM4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm1.176 7.6l-.351-1.2l6.828-2l.351 1.2l-6.828 2z"/>`),
		g.Group(children),
	)
}