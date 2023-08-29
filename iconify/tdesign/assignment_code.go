package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssignmentCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.5 2h6.628A3.496 3.496 0 0 1 12 .5c1.19 0 2.24.594 2.872 1.5H21.5v20h-19V2Zm2 18h15V4h-5.862l-.262-.6a1.5 1.5 0 0 0-2.752 0l-.262.6H4.5v16Zm1.086-8L9.5 8.086L10.914 9.5l-2.5 2.5l2.5 2.5L9.5 15.914L5.586 12ZM14.5 8.086L18.414 12L14.5 15.914L13.086 14.5l2.5-2.5l-2.5-2.5L14.5 8.086Z"/>`),
		g.Group(children),
	)
}