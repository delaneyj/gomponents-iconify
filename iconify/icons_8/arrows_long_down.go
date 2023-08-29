package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsLongDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4v20.063l-4.28-4.282l-1.44 1.407l6 6l.72.72l.72-.72l6-6l-1.44-1.406L17 24.063V4h-2z"/>`),
		g.Group(children),
	)
}