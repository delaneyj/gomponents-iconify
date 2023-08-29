package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fonts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M30.482 27.781h7.521m-7.521 0c-2.034 0-3.683 1.662-3.683 3.712c0 2.05 1.649 3.711 3.683 3.711h1.951c3.07 0 5.56-2.508 5.56-5.602"/><path d="M33.101 20.359h5.497c2.958 0 4.902 1.3 4.902 4.818v10.027h-5.497V25.177c0-3.517-1.944-4.818-4.902-4.818c-2.605 0-3.552.118-4.902 1.47"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.886 27.781H6.959M4.5 35.204l7.423-22.408l7.422 22.408"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.923 12.796h5.337l7.422 22.408h-5.337m11.961-7.423v7.423"/>`),
		g.Group(children),
	)
}