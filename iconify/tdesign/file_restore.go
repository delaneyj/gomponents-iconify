package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileRestore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V12h-2V9h-6V3H5v18h7v2H3V1Zm12 2.414V7h3.586L15 3.414ZM18.414 13l-1 1H18a5 5 0 1 1-5 5v-1h2v1a3 3 0 1 0 3-3h-.586l1 1L17 18.414L13.586 15L17 11.586L18.414 13Z"/>`),
		g.Group(children),
	)
}