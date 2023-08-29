package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaseballEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0zm2.85 7.24l-3-4.85a.86.86 0 0 0-.5-.39H2.5a.5.5 0 0 0 0 1H5l.92 1.09l-2.73 3a.49.49 0 0 0-.19.41a.5.5 0 0 0 .5.5a.49.49 0 0 0 .33-.13l3-2.71L9 10.81a.49.49 0 0 0 .38.19a.5.5 0 0 0 .5-.5a.49.49 0 0 0-.03-.26zM4 .28A.28.28 0 0 0 3.72 0a.49.49 0 0 0-.25.16L2 4.59a.48.48 0 0 0 0 .13c0 .155.125.28.28.28a.49.49 0 0 0 .26-.16L4 .41a.472.472 0 0 0 0-.13z" fill="currentColor"/>`),
		g.Group(children),
	)
}