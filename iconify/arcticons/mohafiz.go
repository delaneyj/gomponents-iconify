package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mohafiz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="10.989" fill="none" stroke="currentColor" stroke-miterlimit="10"/><circle cx="24" cy="24" r="19.049" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m19.869 13.646l.029-5.425l-2.138-1.989m10.371 7.414l-.029-5.425l2.138-1.989"/><circle cx="24" cy="3.726" r="1.226" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m13.646 28.131l-5.425-.029l-1.989 2.138m7.414-10.371l-5.425.029l-1.989-2.138"/><circle cx="3.726" cy="24" r="1.226" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m28.131 34.354l-.029 5.425l2.138 1.989m-10.371-7.414l.029 5.425l-2.138 1.989"/><circle cx="24" cy="44.274" r="1.226" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m34.354 19.869l5.425.029l1.989-2.138m-7.414 10.371l5.425-.029l1.989 2.138"/><circle cx="44.274" cy="24" r="1.226" fill="none" stroke="currentColor" stroke-miterlimit="10"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.117 27.342a7.642 7.642 0 1 1-10.456-10.459"/><path fill="currentColor" d="m25.376 18.233l.978 3.008h3.164l-2.56 1.86l.978 3.009l-2.56-1.86l-2.559 1.86l.978-3.009l-2.56-1.86H24.4l.976-3.008m0-1.05a1.05 1.05 0 0 0-.998.725l-.742 2.284h-2.401a1.05 1.05 0 0 0-.617 1.898l1.942 1.412l-.742 2.283a1.05 1.05 0 0 0 1.615 1.173l1.943-1.411l1.942 1.411a1.05 1.05 0 0 0 1.615-1.173l-.742-2.284l1.943-1.411a1.05 1.05 0 0 0-.617-1.898h-2.401l-.742-2.284a1.05 1.05 0 0 0-.998-.725Z"/>`),
		g.Group(children),
	)
}