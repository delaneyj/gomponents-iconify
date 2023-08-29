package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 1h14v22H5V1Zm2.5 10H17v10H7V11h.5ZM15 6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-4 13v-5h2v5h-2Zm1-13.998a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}