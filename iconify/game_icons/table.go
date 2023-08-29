package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M41 265v30h430v-30H41zm39 48v158.066h32V313H80zm320 0v158.066h32V313h-32z"/>`),
		g.Group(children),
	)
}