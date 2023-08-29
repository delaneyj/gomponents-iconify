package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatOneMinimalisticOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.47 2.47a.75.75 0 0 1 1.06 0l2 2A.75.75 0 0 1 11 5.75H9a6.25 6.25 0 1 0 0 12.5h.5a.75.75 0 0 1 0 1.5H9a7.75 7.75 0 0 1 0-15.5h.19l-.72-.72a.75.75 0 0 1 0-1.06ZM13.75 5a.75.75 0 0 1 .75-.75h.5a7.75 7.75 0 0 1 0 15.5h-.19l.72.72a.75.75 0 1 1-1.06 1.06l-2-2a.75.75 0 0 1 .53-1.28h2a6.25 6.25 0 1 0 0-12.5h-.5a.75.75 0 0 1-.75-.75Zm-6.5 7a.75.75 0 0 1 .75-.75h8a.75.75 0 0 1 0 1.5H8a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}