package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ICertificatePaperOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M18 11a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H19a1 1 0 0 1-1-1Zm-3 5a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H15Zm-1 5a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H15a1 1 0 0 1-1-1Zm1 3a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2H15Z"/><path fill-rule="evenodd" d="M38 36a4 4 0 0 1-4 4h-3v4l-3-1.5l-3 1.5v-4H14a4 4 0 0 1-4-4V8a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v28ZM14 6a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h11v-2.354a4 4 0 1 1 6 0V38h3a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H14Zm15 30.874a4.01 4.01 0 0 1-2 0v3.89l1-.5l1 .5v-3.89ZM28 35a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}