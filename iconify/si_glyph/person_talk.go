package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonTalk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.524 12.285c-.86 0-2.015-.449-2.601-1.186C.114 11.099.029 15 .029 15h10.937c.001 0 .22-3.917-2.937-3.917c-.584.745-1.643 1.202-2.505 1.202zm2.414-6.021c0 1.251-1.105 3.643-2.469 3.643C4.105 9.907 3 7.515 3 6.264C3 5.015 4.104 4 5.469 4c1.364 0 2.469 1.015 2.469 2.264zM12.527.041c-1.91 0-3.461 1.158-3.461 2.59c0 1.306 1.294 2.382 2.971 2.561l-.986 1.812l2.756-1.969c1.277-.381 2.182-1.313 2.182-2.404c-.001-1.432-1.55-2.59-3.462-2.59z"/>`),
		g.Group(children),
	)
}