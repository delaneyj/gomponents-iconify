package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullhorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.5 8.974l6 2.333V2.693l-6 2.333v3.948Zm6.138-7.944L7 4H3a3 3 0 0 0 0 6h.969l1.066 4.445A2.029 2.029 0 0 0 7.008 16c1.524 0 2.504-1.618 1.797-2.969l-1.53-2.924l7.363 2.863A1 1 0 0 0 16 12.038V1.962a1 1 0 0 0-1.362-.932ZM5.51 10h.016l1.95 3.726a.529.529 0 1 1-.983.369L5.511 10ZM3 5.5h4v3H3a1.5 1.5 0 1 1 0-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}