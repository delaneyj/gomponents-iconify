package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SadTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0zm3 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2zM5 4a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm6.002 8.199C10.39 11.181 9.275 10.5 8 10.5s-2.389.681-3.002 1.699l-1.286-.772C4.586 9.973 6.179 9 8 9s3.414.973 4.288 2.427l-1.286.772z"/>`),
		g.Group(children),
	)
}