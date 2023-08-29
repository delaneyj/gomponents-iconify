package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.5 5L5 9v16h5.5v3h3l3-3H21l6-6V5zM9 7h16v11l-3 3h-6l-3 3v-3H9zm5 4v6h2v-6zm5 0v6h2v-6z"/>`),
		g.Group(children),
	)
}