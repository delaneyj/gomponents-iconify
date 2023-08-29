package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12.948 3.21A1 1 0 0 0 12 2.632a1 1 0 0 0-.948.576l-5.917 12.69a1 1 0 1 0 1.813.845l1.45-3.111h7.203l1.451 3.111a1 1 0 0 0 1.813-.845L12.948 3.209Zm1.72 8.422L12 5.909l-2.669 5.723h5.338Z" clip-rule="evenodd"/><path d="M6 19.368a1 1 0 0 0 0 2h12a1 1 0 1 0 0-2H6Z"/></g>`),
		g.Group(children),
	)
}