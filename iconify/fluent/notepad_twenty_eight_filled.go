package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotepadTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2a.75.75 0 0 1 .75.75V4h3.75V2.75a.75.75 0 0 1 1.5 0V4h.75A2.25 2.25 0 0 1 23 6.25v12.246h-4.754a2.25 2.25 0 0 0-2.25 2.25V25.5H6.75a2.25 2.25 0 0 1-2.25-2.25v-17A2.25 2.25 0 0 1 6.75 4H8V2.75a.75.75 0 0 1 1.5 0V4h3.75V2.75A.75.75 0 0 1 14 2Zm-6 8.25c0 .414.336.75.75.75h10a.75.75 0 0 0 0-1.5h-10a.75.75 0 0 0-.75.75Zm0 4.5c0 .414.336.75.75.75h10a.75.75 0 0 0 0-1.5h-10a.75.75 0 0 0-.75.75Zm0 4.5c0 .414.336.75.75.75h4.5a.75.75 0 0 0 0-1.5h-4.5a.75.75 0 0 0-.75.75Zm9.496 5.81l5.065-5.064h-4.315a.75.75 0 0 0-.75.75v4.315Z"/>`),
		g.Group(children),
	)
}