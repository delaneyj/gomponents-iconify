package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationDownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.975 20.625q-.2 0-.4-.075t-.35-.225l-2-2q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l.35.325v-12.2q0-.425.288-.713T6 4.026q.425 0 .713.288t.287.71v12.2l.35-.35q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-2.05 2.05q-.15.15-.338.225t-.387.075ZM11.2 16.65q-.5.2-.85-.05t-.35-.8q0-.25.163-.475t.387-.3l2.25-.775V9.8L10.55 9q-.225-.075-.388-.3T10 8.2q0-.55.35-.8t.85-.05l8.875 3.325q.375.15.65.537t.275.813q0 .425-.275.8t-.65.525l-8.875 3.3Zm3.2-2.975l4.55-1.6v-.1l-4.55-1.6v3.3Z"/>`),
		g.Group(children),
	)
}