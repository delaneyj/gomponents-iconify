package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Edx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.774 12.3H8.605L4.5 27.9h22.168l4.106-15.6ZM43.5 20.1H28.721l-4.105 15.6h14.779L43.5 20.1ZM32.48 24l3.156 8m2.144-8l-7.444 8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.855 23.594c-.632.754-1.58 1.262-2.51 1.262h0c-1.382 0-2.201-1.12-1.831-2.5l.435-1.626c.37-1.382 1.79-2.501 3.171-2.501h0c1.381 0 2.201 1.12 1.83 2.5l-.217.813h-5.002m12.913-.812c.37-1.382-.45-2.501-1.831-2.501h0c-1.381 0-2.8 1.12-3.171 2.5l-.436 1.626c-.37 1.381.45 2.5 1.831 2.5h0c1.381 0 2.8-1.119 3.171-2.5m-.67 2.501l2.68-10.003"/>`),
		g.Group(children),
	)
}