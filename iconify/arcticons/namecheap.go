package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Namecheap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.274 36.644a3.686 3.686 0 0 1-3.183-1.818l-10.54-17.909a3.689 3.689 0 0 1 6.358-3.742l10.54 17.908a3.69 3.69 0 0 1-3.175 5.561Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.27 25.682l7.361-12.507a3.689 3.689 0 0 1 6.359 3.742L32.45 34.826M18.73 22.319l-7.361 12.507a3.689 3.689 0 0 1-6.358-3.743l10.54-17.908"/>`),
		g.Group(children),
	)
}