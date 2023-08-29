package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inmage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 1 2.5 24A21.503 21.503 0 0 1 24 2.5Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M16.517 15.499h14.966a1.018 1.018 0 0 1 1.018 1.018v14.966a1.018 1.018 0 0 1-1.018 1.018H16.517a1.018 1.018 0 0 1-1.018-1.018V16.517a1.018 1.018 0 0 1 1.018-1.018Z"/>`),
		g.Group(children),
	)
}