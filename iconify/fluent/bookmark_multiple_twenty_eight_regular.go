package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMultipleTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.401 3.5A2.999 2.999 0 0 1 11 2h6.25A5.75 5.75 0 0 1 23 7.75V20a3 3 0 0 1-1.5 2.599V7.75a4.25 4.25 0 0 0-4.25-4.25H8.401ZM8 5a3 3 0 0 0-3 3v17.25a.75.75 0 0 0 1.166.624l6.334-4.223l6.334 4.223A.75.75 0 0 0 20 25.25V8a3 3 0 0 0-3-3H8ZM6.5 8A1.5 1.5 0 0 1 8 6.5h9A1.5 1.5 0 0 1 18.5 8v15.849l-5.584-3.723a.75.75 0 0 0-.832 0L6.5 23.849V8Z"/>`),
		g.Group(children),
	)
}