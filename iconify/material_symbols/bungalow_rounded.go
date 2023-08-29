package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BungalowRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14q.425 0 .713-.288T13 13q0-.425-.288-.713T12 12q-.425 0-.713.288T11 13q0 .425.288.713T12 14Zm-4 7q-.425 0-.713-.288T7 20v-5.2l-.575.925q-.225.35-.625.438t-.75-.138Q4.7 15.8 4.6 15.4t.125-.75l6.425-10.3q.15-.25.375-.35T12 3.9q.25 0 .475.1t.375.35l6.425 10.3q.225.35.125.75t-.45.625q-.35.225-.75.125t-.625-.45L17 14.8V20q0 .425-.288.713T16 21h-3v-4q0-.425-.288-.713T12 16q-.425 0-.713.288T11 17v4H8Z"/>`),
		g.Group(children),
	)
}