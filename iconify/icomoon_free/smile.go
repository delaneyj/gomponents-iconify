package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13zM4 5a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm6 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0zm1.002 4.801l1.286.772C11.414 12.027 9.821 13 8 13s-3.413-.973-4.288-2.427l1.286-.772C5.61 10.819 6.725 11.5 8 11.5s2.389-.681 3.002-1.699z"/>`),
		g.Group(children),
	)
}