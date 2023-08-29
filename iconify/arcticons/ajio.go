package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ajio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.092 37.344q-9.965 5.836-18.353-.178l2.539-4.69l13.412.085Zm-4.932-9.796l-8.275-.017l4.073-8.041Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.788 33.443L23.958 8.218l-12.83 25.343C2.51 19.991 13.41 8.218 23.959 8.218c14.314 0 20.291 16.616 12.83 25.225Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 23.946a21.497 21.497 0 1 1-.001-.173"/>`),
		g.Group(children),
	)
}