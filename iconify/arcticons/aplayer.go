package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aplayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.647 30.442a18.161 18.161 0 1 1 0-12.885"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 28.171l-2.711-8.342l-2.816 8.342m.939-2.815h3.65"/><circle cx="22.396" cy="24" r="7.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="22.396" cy="24" r=".782" fill="currentColor"/>`),
		g.Group(children),
	)
}