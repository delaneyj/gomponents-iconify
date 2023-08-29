package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M160 160v704h704V160H160zm-32-64h768a32 32 0 0 1 32 32v768a32 32 0 0 1-32 32H128a32 32 0 0 1-32-32V128a32 32 0 0 1 32-32z"/><path fill="currentColor" d="M384 288q64 0 64 64t-64 64q-64 0-64-64t64-64zM185.408 876.992l-50.816-38.912L350.72 556.032a96 96 0 0 1 134.592-17.856l1.856 1.472l122.88 99.136a32 32 0 0 0 44.992-4.864l216-269.888l49.92 39.936l-215.808 269.824l-.256.32a96 96 0 0 1-135.04 14.464l-122.88-99.072l-.64-.512a32 32 0 0 0-44.8 5.952L185.408 876.992z"/>`),
		g.Group(children),
	)
}