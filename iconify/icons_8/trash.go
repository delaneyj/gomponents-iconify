package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4c-.522 0-1.06.185-1.438.563C13.185 4.94 13 5.478 13 6v1H7v2h1v16c0 1.645 1.355 3 3 3h12c1.645 0 3-1.355 3-3V9h1V7h-6V6c0-.522-.185-1.06-.563-1.438C20.06 4.186 19.523 4 19 4h-4zm0 2h4v1h-4V6zm-5 3h14v16c0 .555-.445 1-1 1H11c-.555 0-1-.445-1-1V9zm2 3v11h2V12h-2zm4 0v11h2V12h-2zm4 0v11h2V12h-2z"/>`),
		g.Group(children),
	)
}