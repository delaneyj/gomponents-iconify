package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 30h-2V20H12v10h-2V20a2.002 2.002 0 0 1 2-2h12a2.002 2.002 0 0 1 2 2zM5.17 16L2 19.17l1.411 1.419L8 16l-4.58-4.58L2 12.83L5.17 16zM24 14H12a2.002 2.002 0 0 1-2-2V2h2v10h12V2h2v10a2.002 2.002 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}