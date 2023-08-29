package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v11H4v7h9v2H2V2Zm2 9h16V4H4v7Zm1.998-4.5h2.004v2.004H5.998V6.5ZM17.75 16a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5ZM14 17.75a3.75 3.75 0 1 1 7.013 1.849l1.651 1.651l-1.414 1.414l-1.651-1.65A3.75 3.75 0 0 1 14 17.75ZM5.998 15.5h2.004v2.004H5.998V15.5Z"/>`),
		g.Group(children),
	)
}