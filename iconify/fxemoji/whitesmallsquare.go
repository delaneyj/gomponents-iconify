package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whitesmallsquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#D9DDE8" d="M371 154.504c0-7.427-6.077-13.504-13.504-13.504H152.82c-6.684 0-12.154 5.469-12.154 12.154V359.18c0 6.684 5.469 12.154 12.154 12.154h204.676c7.427 0 13.504-6.077 13.504-13.504V154.504z"/>`),
		g.Group(children),
	)
}