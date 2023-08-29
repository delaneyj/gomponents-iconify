package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M191.5 376h97A111.627 111.627 0 0 0 400 264.5V248h24a72 72 0 0 0 0-144H80v160.5A111.627 111.627 0 0 0 191.5 376ZM400 136h24a40 40 0 0 1 0 80h-24Zm-288 0h256v128.5a79.589 79.589 0 0 1-79.5 79.5h-97a79.589 79.589 0 0 1-79.5-79.5ZM16 416h480v32H16z"/>`),
		g.Group(children),
	)
}