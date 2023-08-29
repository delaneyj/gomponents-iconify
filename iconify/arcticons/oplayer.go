package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oplayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.905 34.248a21.5 21.5 0 1 1-.01-20.516"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.7 24l-13-7.51v15.02l13-7.51z"/><circle cx="42.905" cy="19.927" r="2.243" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}