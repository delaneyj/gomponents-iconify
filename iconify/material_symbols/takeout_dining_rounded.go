package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TakeoutDiningRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.8 20q-.8 0-1.375-.525T5.8 18.15l-.5-6.6h13.4l-.5 6.6q-.05.8-.625 1.325T16.2 20H7.8ZM4.625 10l-1.9-1.85q-.275-.275-.313-.688T2.7 6.75q.275-.275.7-.275t.7.275l.9.9l-.05-.6l3.475-3.475q.275-.275.638-.425T9.825 3h4.35q.4 0 .763.15t.637.425L19.05 7.05l-.05.6l.9-.9q.275-.275.663-.35t.737.35q.275.325.275.713t-.3.687l-1.9 1.85H4.625Z"/>`),
		g.Group(children),
	)
}