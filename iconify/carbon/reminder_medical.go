package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReminderMedical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 12v2h1v4a3 3 0 0 1-6 0v-4h1v-2h-3v6a5.008 5.008 0 0 0 4 4.899V24a4 4 0 0 1-8 0v-2.184a3 3 0 1 0-2 0V24a6 6 0 0 0 12 0v-1.101A5.008 5.008 0 0 0 30 18v-6Zm-12 6a1 1 0 1 1-1 1a1 1 0 0 1 1-1Z"/><path fill="currentColor" d="M26 4h-4V2h-2v2h-8V2h-2v2H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h4v-2H6V6h4v2h2V6h8v2h2V6h4v4h2V6a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}