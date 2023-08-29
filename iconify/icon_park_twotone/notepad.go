package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTNotepad0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M18 8h-7a1 1 0 0 0-1 1v34a1 1 0 0 0 1 1h28a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-7"/><path fill="#555" d="M18 13V8h3.95a.05.05 0 0 0 .05-.05V6a3 3 0 1 1 6 0v1.95c0 .028.022.05.05.05H32v5a1 1 0 0 1-1 1H19a1 1 0 0 1-1-1Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTNotepad0)"/>`),
		g.Group(children),
	)
}