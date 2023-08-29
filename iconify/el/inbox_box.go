package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m93.75 637.5l277.837-375h456.825l277.838 375H900l-150 150H450l-150-150H93.75zm235.538-450L0 637.5v375h1200v-375l-329.287-450H329.288z"/>`),
		g.Group(children),
	)
}