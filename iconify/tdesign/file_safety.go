package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSafety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11.5h-2V9h-6V3H5v18h6.5v2H3V1Zm12 2.414V7h3.586L15 3.414ZM13.498 13.5h9v5.634a3 3 0 0 1-1.36 2.511l-3.14 2.052l-3.14-2.052a3 3 0 0 1-1.36-2.511V13.5Zm2 2v3.634a1 1 0 0 0 .453.837l2.047 1.337l2.047-1.337a1 1 0 0 0 .453-.837V15.5h-5Z"/>`),
		g.Group(children),
	)
}