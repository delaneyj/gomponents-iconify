package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shareable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.8 10a3.2 3.2 0 1 0 6.401 0A3.2 3.2 0 0 0 6.8 10zM4.529 8.8a5.6 5.6 0 0 1 9.43-2.76a1.2 1.2 0 1 0 1.697-1.697A8.002 8.002 0 0 0 2.367 7.601H0V10h3.199c.999 0 1.245-.813 1.33-1.2zM16.8 10c-.999 0-1.245.814-1.329 1.199a5.6 5.6 0 0 1-9.43 2.759a1.2 1.2 0 0 0-1.698 1.697A7.972 7.972 0 0 0 10 18a8.005 8.005 0 0 0 7.633-5.6H20V10h-3.2z"/>`),
		g.Group(children),
	)
}