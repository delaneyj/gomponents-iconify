package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeTypeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.3 19.3q-.275.275-.7.275t-.7-.275l-4.3-4.325q-.275-.275-.437-.637T11 13.575v-6.7L9.125 8.75q-.3.3-.713.3t-.712-.3q-.3-.3-.3-.7t.3-.7l3.6-3.625q.15-.15.325-.212T12 3.45q.2 0 .375.063t.325.212L16.3 7.3q.3.3.3.713t-.3.712q-.3.3-.713.313t-.712-.288L13 6.875V13.6l4.3 4.3q.275.275.275.7t-.275.7Zm-10.6.025q-.3-.3-.288-.713t.288-.687l1.775-1.8q.3-.3.713-.3t.712.3q.3.3.288.713t-.313.712L8.1 19.325q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}