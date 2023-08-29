package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Revolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21 12.368v9.68h6.396c2.188 0 4.698-.799 4.698-4.586s-2.388-5.094-4.698-5.094H21ZM5.83 4.5l5.83 7.811V43.5h9.51V28.783h2.037L31.473 43.5H42.17l-9.113-15.453c4.054-1.534 8.345-3.677 8.547-11.717C41.805 8.29 34.69 4.5 29.377 4.5H5.83Z"/>`),
		g.Group(children),
	)
}