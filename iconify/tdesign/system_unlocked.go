package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SystemUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1h22v10h-2V3H3v13h10v2H1V1Zm18.502 13.5c-.69 0-1.25.56-1.25 1.25v.75H24V23h-9v-6.5h1.252v-.75a3.25 3.25 0 0 1 5.54-2.305l.71.705l-1.41 1.418l-.71-.705a1.243 1.243 0 0 0-.88-.363ZM17 18.5V21h5v-2.5h-5ZM2.25 21H13v2H2.25v-2Z"/>`),
		g.Group(children),
	)
}