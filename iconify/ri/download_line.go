package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19h18v2H3v-2Zm10-5.828L19.071 7.1l1.414 1.415L12 17L3.515 8.515L4.929 7.1L11 13.172V2h2v11.172Z"/>`),
		g.Group(children),
	)
}