package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotepadLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M166 128a6 6 0 0 1-6 6H96a6 6 0 0 1 0-12h64a6 6 0 0 1 6 6Zm-6 26H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12Zm54-106v152a30 30 0 0 1-30 30H72a30 30 0 0 1-30-30V48a14 14 0 0 1 14-14h18V24a6 6 0 0 1 12 0v10h36V24a6 6 0 0 1 12 0v10h36V24a6 6 0 0 1 12 0v10h18a14 14 0 0 1 14 14Zm-12 0a2 2 0 0 0-2-2h-18v10a6 6 0 0 1-12 0V46h-36v10a6 6 0 0 1-12 0V46H86v10a6 6 0 0 1-12 0V46H56a2 2 0 0 0-2 2v152a18 18 0 0 0 18 18h112a18 18 0 0 0 18-18Z"/>`),
		g.Group(children),
	)
}