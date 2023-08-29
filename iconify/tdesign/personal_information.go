package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonalInformation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm2 2v18h14V7.414L14.586 3H5Zm7 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-3.5 1.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0ZM6 19a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v1h-2v-1a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1H6v-1Z"/>`),
		g.Group(children),
	)
}