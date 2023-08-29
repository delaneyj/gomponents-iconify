package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BehanceAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.108 15.032a.997.997 0 0 0-1.215.722A2.998 2.998 0 0 1 14.992 15v-1h7a1 1 0 0 0 1-1a5 5 0 1 0-10 0v2a4.998 4.998 0 0 0 9.839 1.247a.999.999 0 0 0-.723-1.215ZM17.992 10a3.011 3.011 0 0 1 2.118.873A3.044 3.044 0 0 1 20.809 12h-5.633a2.995 2.995 0 0 1 2.816-2Zm-2-3h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2ZM9.91 11.718A3.987 3.987 0 0 0 6.992 5h-5a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h5.5a4.492 4.492 0 0 0 2.418-8.282ZM2.992 7h4a2 2 0 1 1 0 4h-4Zm4.5 11h-4.5v-5h4.5a2.5 2.5 0 0 1 0 5Z"/>`),
		g.Group(children),
	)
}