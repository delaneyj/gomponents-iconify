package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicTuneSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7.33a1 1 0 0 0 1-1v-.55l6-1.5v2.94L14.7 8.3a1 1 0 0 0 .24 2h.24L20.24 9h.07l.19-.09l.15-.1a.93.93 0 0 0 .13-.15a.78.78 0 0 0 .1-.15a.55.55 0 0 0 .06-.18a.65.65 0 0 0 0-.19A.24.24 0 0 0 21 8V3a1 1 0 0 0-1.24-1l-8 2A1 1 0 0 0 11 5v1.33a1 1 0 0 0 1 1Zm9.71 13l-9-9l-9-9a1 1 0 0 0-1.42 1.38l8.71 8.7v2.94A3.45 3.45 0 0 0 9.5 15a3.5 3.5 0 1 0 3.5 3.5v-4.09l7.29 7.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM9.5 20a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 9.5 20Z"/>`),
		g.Group(children),
	)
}