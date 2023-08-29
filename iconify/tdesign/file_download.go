package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11.5h-2V9h-6V3H5v18h8v2H3V1Zm12 2.414V7h3.586L15 3.414ZM20 13v7.11l2.508-2.48l1.406 1.422L19 23.91l-4.914-4.858l1.406-1.422L18 20.11V13h2Z"/>`),
		g.Group(children),
	)
}