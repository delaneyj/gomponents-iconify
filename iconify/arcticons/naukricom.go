package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Naukricom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.328 4.5c-16.283 9.06-25.671 25.606 3.283 39c-16.545-.92-27.65-11.046-25.869-21.798C16.482 11.197 24.328 6.864 37.328 4.5Z"/><circle cx="11.657" cy="8.768" r="4.268" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}