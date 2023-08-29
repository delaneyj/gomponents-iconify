package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.03 2.443h-3.642L12.013 6.4L9.63 2.444l-2.646-.001H.831L12.03 21.558L23.168 2.443Zm-6.005 15.15L4.322 4.443h2.824l4.885 8.406l4.847-8.407h2.81Z"/>`),
		g.Group(children),
	)
}