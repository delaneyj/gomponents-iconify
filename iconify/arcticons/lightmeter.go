package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightmeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.13 42.5a33.072 33.072 0 0 0 2.052-11c0-11.87-6.783-26-18.182-26S5.818 19.63 5.818 31.5a33.072 33.072 0 0 0 2.053 11Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.462 10.072a.639.639 0 0 0-.627.531l-.475 3.36a9.376 9.376 0 0 0-2.142 1.248l-3.157-1.274a.634.634 0 0 0-.772.272l-2.535 4.392a.635.635 0 0 0 .152.811l2.674 2.099a8.677 8.677 0 0 0 0 2.47l-2.674 2.099a.632.632 0 0 0-.152.811l2.535 4.391a.626.626 0 0 0 .772.273l3.157-1.274a9.617 9.617 0 0 0 2.142 1.248l.475 3.359a.64.64 0 0 0 .627.533h5.07a.64.64 0 0 0 .627-.533l.475-3.359a9.379 9.379 0 0 0 2.143-1.248l3.156 1.274a.634.634 0 0 0 .773-.273l2.535-4.391a.634.634 0 0 0-.152-.811l-2.675-2.098a8.562 8.562 0 0 0 0-2.471l2.68-2.099a.632.632 0 0 0 .152-.81l-2.534-4.393a.627.627 0 0 0-.773-.272l-3.156 1.274a9.614 9.614 0 0 0-2.142-1.248l-.475-3.36a.652.652 0 0 0-.634-.531Z"/><circle cx="24" cy="22.746" r="7.064" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}