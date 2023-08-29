package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentSaveTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2v6a2 2 0 0 0 2 2h6v10a2 2 0 0 1-2 2h-5.05c.033-.162.05-.329.05-.5v-6.879a2.5 2.5 0 0 0-.732-1.767l-1.122-1.122A2.5 2.5 0 0 0 9.38 11H4V4a2 2 0 0 1 2-2h6Zm1.5.5V8a.5.5 0 0 0 .5.5h5.5l-6-6ZM5 12h3v2H5v-2Zm-2.5 0H4v2.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V12h.379a1.5 1.5 0 0 1 1.06.44l1.122 1.12A1.5 1.5 0 0 1 12 14.622V21.5a1.5 1.5 0 0 1-1.5 1.5H10v-5.5a.5.5 0 0 0-.5-.5h-6a.5.5 0 0 0-.5.5V23h-.5A1.5 1.5 0 0 1 1 21.5v-8A1.5 1.5 0 0 1 2.5 12ZM9 18v5H4v-5h5Z"/>`),
		g.Group(children),
	)
}