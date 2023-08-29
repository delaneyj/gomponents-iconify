package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M22.5 93v326h467V93zm15.1 169.44a6.6 6.6 0 1 1 6.6-6.6a6.6 6.6 0 0 1-6.6 6.6zM427.5 401h-377V109h377zm29-133.32a11.85 11.85 0 1 1 11.85-11.85a11.85 11.85 0 0 1-11.85 11.85z"/>`),
		g.Group(children),
	)
}