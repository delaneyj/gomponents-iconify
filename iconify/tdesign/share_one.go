package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 1.85L23.553 12L11.5 22.15v-6.227c-4.194.2-7.073 1.7-9.186 4.658L.52 19.804c.523-2.617 1.58-5.295 3.478-7.462c1.761-2.014 4.209-3.543 7.502-4.187V1.851Zm2 4.3v3.717l-.858.123c-3.27.467-5.551 1.853-7.14 3.668a11.98 11.98 0 0 0-1.744 2.66C6.096 14.665 8.978 13.9 12.5 13.9h1v3.95L20.448 12L13.5 6.15Z"/>`),
		g.Group(children),
	)
}