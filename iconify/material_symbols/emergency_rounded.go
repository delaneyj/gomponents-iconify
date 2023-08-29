package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmergencyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.25 21q-.425 0-.713-.288T10.25 20v-4.95l-4.3 2.475q-.35.2-.763.1t-.612-.475l-.75-1.3q-.2-.35-.1-.75t.475-.6L8.5 12L4.225 9.525q-.375-.2-.475-.612t.1-.763l.75-1.3q.2-.35.6-.45t.775.1l4.275 2.475V4q0-.425.288-.713T11.25 3h1.5q.425 0 .713.288T13.75 4v4.975l4.3-2.475q.35-.225.763-.112t.612.487l.725 1.275q.2.35.088.763t-.463.612L15.5 12l4.275 2.5q.375.225.475.625t-.1.75l-.75 1.3q-.2.35-.612.45t-.763-.1L13.75 15.05V20q0 .425-.287.713T12.75 21h-1.5Z"/>`),
		g.Group(children),
	)
}