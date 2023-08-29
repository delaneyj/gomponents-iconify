package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 14h-6a1 1 0 1 1 0-2h6a1 1 0 1 1 0 2zm-7-5c0 1.381-.56 2.631-1.464 3.535C11.631 13.44 10.381 14 9 14s-2.631-.56-3.536-1.465C4.56 11.631 4 10.381 4 9s.56-2.631 1.464-3.535C6.369 4.56 7.619 4 9 4s2.631.56 3.536 1.465A4.984 4.984 0 0 1 14 9zm-5 6c-3.75 0-6 2-6 4c0 1 2.25 2 6 2c3.518 0 6-1 6-2c0-2-2.354-4-6-4z"/>`),
		g.Group(children),
	)
}