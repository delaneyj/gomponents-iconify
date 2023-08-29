package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoilMoisture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.5 28a5.385 5.385 0 0 1-5.5-5.249a5.384 5.384 0 0 1 .874-2.83l3.616-5.383a1.217 1.217 0 0 1 2.02 0l3.55 5.277a5.492 5.492 0 0 1 .94 2.936A5.385 5.385 0 0 1 24.5 28Zm0-11.38l-2.936 4.367A3.359 3.359 0 0 0 21 22.751a3.51 3.51 0 0 0 7 0a3.436 3.436 0 0 0-.63-1.867Z"/><circle cx="5" cy="13" r="1" fill="currentColor"/><circle cx="11" cy="19" r="1" fill="currentColor"/><circle cx="15" cy="25" r="1" fill="currentColor"/><circle cx="17" cy="15" r="1" fill="currentColor"/><circle cx="13" cy="11" r="1" fill="currentColor"/><circle cx="27" cy="11" r="1" fill="currentColor"/><circle cx="9" cy="27" r="1" fill="currentColor"/><circle cx="3" cy="21" r="1" fill="currentColor"/><path fill="currentColor" d="M2 6h28v2H2z"/>`),
		g.Group(children),
	)
}