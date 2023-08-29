package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Droplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14v.5M12 4c-1.262 1.683-3.055 3.896-4.708 5.896c-2.288 2.767-1.796 6.907 1.115 9.009v0a6.137 6.137 0 0 0 7.186 0v0c2.91-2.102 3.403-6.242 1.116-9.009C15.055 7.896 13.262 5.683 12 4Z"/>`),
		g.Group(children),
	)
}