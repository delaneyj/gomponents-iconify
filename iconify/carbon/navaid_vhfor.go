package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavaidVhfor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 28H10a1 1 0 0 1-.844-.463l-7-11a1 1 0 0 1 0-1.074l7-11A1 1 0 0 1 10 4h12a1 1 0 0 1 .844.463l7 11a1 1 0 0 1 0 1.074l-7 11A1 1 0 0 1 22 28Zm-11.451-2H21.45l6.363-10L21.45 6h-10.9L4.186 16Z"/><path fill="currentColor" d="M19.5 24h-7a1 1 0 0 1-.841-.46l-4.5-7a1.002 1.002 0 0 1 0-1.08l4.5-7A1 1 0 0 1 12.5 8h7a1 1 0 0 1 .841.46l4.5 7a1.002 1.002 0 0 1 0 1.08l-4.5 7a1 1 0 0 1-.841.46Zm-6.454-2h5.908l3.857-6l-3.857-6h-5.908l-3.857 6Z"/>`),
		g.Group(children),
	)
}