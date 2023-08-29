package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm15.586 6L15 3.414V7h3.586ZM13 3H5v18h14V9h-6V3Zm-1 7h4v2h-2v4a2.5 2.5 0 1 1-2-2.45V10Zm0 6a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0Z"/>`),
		g.Group(children),
	)
}