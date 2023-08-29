package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemMarked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v10h-2V3H3v13h9.5v2H1V1Zm13.5 11.996H23v10.295l-4.247-2.617L14.5 23.29V12.996Zm2 2v4.715l2.254-1.385L21 19.709v-4.713h-4.5ZM2.25 21H12.5v2H2.25v-2Z"/>`),
		g.Group(children),
	)
}