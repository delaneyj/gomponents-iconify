package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShrinkScreenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<rect width="15" height="10" x="15" y="18" fill="currentColor" rx="2"/><path fill="currentColor" d="M12 10v3.586L7.707 9.293l-1.414 1.414L10.586 15H7v2h7v-7h-2z"/><path fill="currentColor" d="M13 22H4a2.002 2.002 0 0 1-2-2V7a2.002 2.002 0 0 1 2-2h22a2.002 2.002 0 0 1 2 2v9h-2V7H4v13h9Z"/>`),
		g.Group(children),
	)
}