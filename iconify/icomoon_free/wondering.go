package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wondering(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13zm3.652 7.9l.351 1.2l-6.828 2l-.351-1.2zM4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0z"/>`),
		g.Group(children),
	)
}