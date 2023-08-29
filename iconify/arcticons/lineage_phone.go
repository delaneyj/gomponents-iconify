package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineagePhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.143 24a34.368 34.368 0 0 1 6.03-19.5H19.88a34.54 34.54 0 0 0 0 39h8.294a34.368 34.368 0 0 1-6.03-19.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.87 16.836h6.096a3.109 3.109 0 0 0 2.536-1.311l2.003-2.826a3.512 3.512 0 0 0 .647-2.031h0a3.512 3.512 0 0 0-.647-2.031l-2.45-3.457a1.613 1.613 0 0 0-1.316-.68h-1.566M22.87 31.164h6.096a3.109 3.109 0 0 1 2.536 1.311l2.003 2.826a3.512 3.512 0 0 1 .647 2.031h0a3.512 3.512 0 0 1-.647 2.031l-2.45 3.457a1.613 1.613 0 0 1-1.316.68h-1.566"/>`),
		g.Group(children),
	)
}