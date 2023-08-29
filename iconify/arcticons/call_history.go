package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallHistory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.595 15a18.898 18.898 0 0 1 1.943-2.92c.734-.909.628-2.236-.198-3.062l-3.832-3.832c-.985-.985-2.595-.893-3.49.175c-9.015 10.762-9.015 26.515 0 37.277c.895 1.068 2.502 1.163 3.487.177l3.417-3.416c1.247-1.247 1.35-2.572.616-3.48a18.894 18.894 0 0 1-1.943-2.921a3.846 3.846 0 0 0-3.392-2.011h-3.28a24.528 24.528 0 0 1 0-13.975h3.28a3.849 3.849 0 0 0 3.392-2.011Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.378 24.698l-2.88 2.863l2.88 2.863m-2.88-2.863h13.245m-2.88-5.033l2.88-2.863l-2.88-2.863m2.88 2.863H23.498"/>`),
		g.Group(children),
	)
}