package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebcamBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M168 104a40 40 0 1 0-40 40a40 40 0 0 0 40-40Zm-56 0a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm112 92h-84v-8.87a84 84 0 1 0-24 0V196H32a12 12 0 0 0 0 24h192a12 12 0 0 0 0-24ZM68 104a60 60 0 1 1 60 60a60.07 60.07 0 0 1-60-60Z"/>`),
		g.Group(children),
	)
}