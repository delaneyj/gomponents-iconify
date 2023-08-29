package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemLocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v10h-2V3H3v13h10v2H1V1Zm18.502 13.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25Zm3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75H15V23h9v-6.5h-1.248Zm-.752 2V21h-5v-2.5h5ZM2.25 21H13v2H2.25v-2Z"/>`),
		g.Group(children),
	)
}