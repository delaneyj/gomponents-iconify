package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmReel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M43 39V24h-4v15c0 5 4 9 9 9v-4c-2.8 0-5-2.2-5-5z"/><circle cx="24" cy="24" r="19" fill="#90A4AE"/><circle cx="24" cy="24" r="2" fill="#37474F"/><g fill="#253278"><circle cx="24" cy="14" r="5"/><circle cx="24" cy="34" r="5"/><circle cx="34" cy="24" r="5"/><circle cx="14" cy="24" r="5"/></g>`),
		g.Group(children),
	)
}