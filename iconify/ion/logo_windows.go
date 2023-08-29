package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWindows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M480 265H232v179l248 36V265Zm-264 0H32v150l184 26.7V265ZM480 32L232 67.4V249h248V32ZM216 69.7L32 96v153h184V69.7Z"/>`),
		g.Group(children),
	)
}