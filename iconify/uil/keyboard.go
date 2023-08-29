package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.21 13.29a.93.93 0 0 0-.33-.21a1 1 0 0 0-.76 0a.9.9 0 0 0-.54.54a1 1 0 1 0 1.84 0a1 1 0 0 0-.21-.33ZM13.5 11h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2Zm-4 0h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2Zm-3-2h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2ZM20 5H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3Zm1 11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-6-3H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Zm3.5-4h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm.71 4.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a.93.93 0 0 0-.33.21a1 1 0 0 0-.21.33a1 1 0 1 0 1.92.38a.84.84 0 0 0-.08-.38a1 1 0 0 0-.21-.33Z"/>`),
		g.Group(children),
	)
}