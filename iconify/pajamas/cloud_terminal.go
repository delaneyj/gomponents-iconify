package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudTerminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.256 6.041a3.75 3.75 0 0 1 7.348-.832l.152.528l.55.014a2.25 2.25 0 0 1 1.069 4.198a.75.75 0 1 0 .75 1.299a3.75 3.75 0 0 0-1.25-6.946a5.251 5.251 0 0 0-10.035.974a3.25 3.25 0 0 0-.896 6.2a.75.75 0 1 0 .603-1.373A1.75 1.75 0 0 1 3.25 6.75h.967l.039-.709ZM6.22 7.22a.75.75 0 0 1 1.06 0l1.75 1.75l.53.53l-.53.53l-1.75 1.75a.75.75 0 0 1-1.06-1.06L7.44 9.5L6.22 8.28a.75.75 0 0 1 0-1.06ZM8 13.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 0 1.5h-2.5a.75.75 0 0 1-.75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}