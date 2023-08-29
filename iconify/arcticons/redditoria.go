package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redditoria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.197 4.5h11.074a9.126 9.126 0 0 1 9.127 9.126v0a9.126 9.126 0 0 1-9.127 9.127H13.197h0V5.5a1 1 0 0 1 1-1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.197 22.753V42.5a1 1 0 0 0 1 1h18.715a1.887 1.887 0 0 0 1.36-3.195l-16.88-17.552"/>`),
		g.Group(children),
	)
}