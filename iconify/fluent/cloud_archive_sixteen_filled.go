package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudArchiveSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.03 5.507a4 4 0 0 1 7.94 0A3.248 3.248 0 0 1 14.49 7H8a2 2 0 0 0-1 3.732V12H4.25a3.25 3.25 0 0 1-.22-6.493ZM8 8a1 1 0 0 0 0 2h7a1 1 0 1 0 0-2H8Zm7 3H8v3a2 2 0 0 0 2 2h3a2 2 0 0 0 2-2v-3Zm-5 1.5a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}