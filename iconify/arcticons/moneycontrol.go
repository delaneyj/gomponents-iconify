package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moneycontrol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.617 22.091c0-2.096 1.606-3.796 3.587-3.796h0c1.98 0 3.586 1.7 3.586 3.796v6.264m-7.173-10.059v10.059m7.173-6.264c0-2.096 1.605-3.796 3.586-3.796h0c1.98 0 3.586 1.7 3.586 3.796v6.264"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.379 14.536V34.99c0 .744.566 1.138 1.276.919c2.724-.839 9.468-2.536 16.93-1.559c.737.097 1.914.355 2.64.518c9.584 2.158 14.933-.032 17.06-1.31c.637-.383 1.093-1.385 1.093-2.129V12.755c0-.743-.566-1.139-1.279-.928c-2.733.808-9.51 2.37-17.061.785c-.728-.153-1.899-.44-2.629-.58c-8.354-1.606-14.314-.154-16.775.673c-.705.237-1.255 1.087-1.255 1.83Z"/>`),
		g.Group(children),
	)
}