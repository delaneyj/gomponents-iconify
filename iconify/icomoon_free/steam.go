package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Steam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 4.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 11 4.5zm3.975-2.475A3.5 3.5 0 0 0 9.036 4L5.838 8.847c-.43.022-.856.132-1.249.328L2.122 7.247A1.313 1.313 0 0 0 .506 9.315l2.436 1.905a3.157 3.157 0 1 0 6.168.228l3.891-3.484a3.5 3.5 0 0 0 1.975-5.939zM6 14.105A2.105 2.105 0 0 1 3.895 12l.001-.033l1.046.817a1.31 1.31 0 1 0 1.616-2.068l-.992-.776A2.105 2.105 0 1 1 6 14.105zM12.5 7a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z"/>`),
		g.Group(children),
	)
}