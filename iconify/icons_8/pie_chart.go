package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm-1.125 2.063c.04-.005.084.003.125 0v11.343l.28.313l7.783 7.75C21.154 26.06 18.686 27 16 27C9.913 27 5 22.087 5 16c0-5.707 4.32-10.375 9.875-10.938zm2.125 0A10.956 10.956 0 0 1 26.938 15H17V5.062zM18.438 17h8.5c-.207 2.294-1.077 4.395-2.47 6.063L18.44 17z"/>`),
		g.Group(children),
	)
}