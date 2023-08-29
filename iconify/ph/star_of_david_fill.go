package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOfDavidFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M201.21 128L231 76a8 8 0 0 0-7-12h-59.35L135 12a8 8 0 0 0-13.9 0L91.33 64H32a8 8 0 0 0-6.95 12l29.72 52l-29.72 52a8 8 0 0 0 7 12h59.28l29.72 52a8 8 0 0 0 13.9 0l29.7-52H224a8 8 0 0 0 7-12Zm-18.42 0l-27.42 48h-54.75L73.2 128l27.42-48h54.75Z"/>`),
		g.Group(children),
	)
}