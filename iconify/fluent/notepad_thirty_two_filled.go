package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotepadThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 4a1 1 0 1 0-2 0v1H9a3 3 0 0 0-3 3v18a3 3 0 0 0 3 3h8.5v-5.5A2.5 2.5 0 0 1 20 21h6V8a3 3 0 0 0-3-3h-1V4a1 1 0 1 0-2 0v1h-3V4a1 1 0 1 0-2 0v1h-3V4Zm0 7h8a1 1 0 1 1 0 2h-8a1 1 0 1 1 0-2Zm-1 6a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1Zm1 4h3a1 1 0 1 1 0 2h-3a1 1 0 1 1 0-2Zm8.25 7.212a3 3 0 0 1-.75.502V23.5a.5.5 0 0 1 .5-.5h5.664a3 3 0 0 1-.637.832l-4.778 4.38Z"/>`),
		g.Group(children),
	)
}