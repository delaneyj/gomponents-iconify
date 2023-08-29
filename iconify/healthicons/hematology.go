package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hematology(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm15.018 28.993c4.48-.003 7.997-3.463 7.994-7.863C32.008 23.044 24 12.993 24 12.993s-7.992 9.75-7.988 14.15c.003 4.4 3.526 7.854 8.006 7.85Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}