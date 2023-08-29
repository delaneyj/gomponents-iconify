package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adventures(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm-1.125 2.063c.04-.005.084.003.125 0V6h2v-.938A10.958 10.958 0 0 1 26.938 15H26v2h.938A10.958 10.958 0 0 1 17 26.938V26h-2v.938A10.958 10.958 0 0 1 5.062 17H6v-2h-.938c.472-5.243 4.587-9.41 9.813-9.938zm7.22 4.843l-8.5 3.688l-3.69 8.5l8.5-3.688l3.69-8.5zM16 14.5c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5z"/>`),
		g.Group(children),
	)
}