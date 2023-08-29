package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeBuf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g stroke-width=".092"><path fill="#77e1ff" d="M4.03 2.034H2.033v21.947h1.995zm3.99 0H6.024v21.947H8.02zm3.99 0h-1.995v21.947h1.995zm3.99 0h-1.995v21.947H16zm3.99 0h-1.995v21.947h1.996zm3.99 0h-1.995v21.947h1.996z"/><path fill="#161ede" d="M10.015 8.02H8.02v21.947h1.995zm3.991 0H12.01v21.947h1.996zm3.989 0H16v21.947h1.995zm3.99 0H19.99v21.947h1.995zm3.991 0h-1.995v21.947h1.995zm3.99 0h-1.995v21.947h1.995z"/></g>`),
		g.Group(children),
	)
}