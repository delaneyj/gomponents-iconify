package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorPasteWord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.38 2L15 5v1H5V5l2.64-3h4.74zM10 5c.55 0 1-.45 1-1s-.45-1-1-1s-1 .45-1 1s.45 1 1 1zm8 12V5c0-.55-.45-1-1-1h-1.54l.54.63V7H4V4.62L4.55 4H3c-.55 0-1 .45-1 1v12c0 .55.45 1 1 1h14c.55 0 1-.45 1-1zm-3-8l-2 7h-2l-1-5l-1 5H6.92L5 9h2l1 5l1-5h2l1 5l1-5h2z"/>`),
		g.Group(children),
	)
}