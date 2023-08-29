package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatchFixes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.48 18.71a3.996 3.996 0 0 1-5.163-5.272l2.619 2.619l2.12-2.121l-2.618-2.619a3.988 3.988 0 0 1 5.2 5.308l1.933 1.933A7.96 7.96 0 0 0 20 14A17.115 17.115 0 0 0 13.5.67a21.494 21.494 0 0 1 .74 4.8a3.47 3.47 0 0 1-3.41 3.73A3.64 3.64 0 0 1 7.2 5.47l.03-.36A13.768 13.768 0 0 0 4 14a8 8 0 0 0 12.43 6.66Z"/>`),
		g.Group(children),
	)
}