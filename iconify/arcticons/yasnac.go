package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yasnac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.296 18.865l-9.403 9.404l-5.197-5.304"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.381 24.026V10.364L24 4.5l15.619 5.864v13.662c-.479 9-7.588 17.54-15.619 19.474c-8.031-1.935-15.14-10.473-15.619-19.474Z"/>`),
		g.Group(children),
	)
}