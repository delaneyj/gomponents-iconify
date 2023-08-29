package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 64H237q-10 0-17-9l-19-27Q183 0 147 0H64Q37 0 18.5 18.5T0 64v256q0 27 18.5 45.5T64 384h384q27 0 45.5-18.5T512 320V128q0-27-18.5-45.5T448 64zm21 256q0 21-21 21H64q-21 0-21-21V64q0-21 21-21h83q9 0 17 8l17 26q21 30 56 30h211q21 0 21 21v192z"/>`),
		g.Group(children),
	)
}