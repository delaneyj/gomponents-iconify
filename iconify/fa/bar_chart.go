package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M640 768v512H384V768h256zm384-512v1024H768V256h256zm1024 1152v128H0V0h128v1408h1920zm-640-896v768h-256V512h256zm384-384v1152h-256V128h256z"/>`),
		g.Group(children),
	)
}