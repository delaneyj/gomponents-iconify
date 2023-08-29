package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeckOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.425 0-.713-.288T11 21V9H3.575Q3.2 9 3.1 8.65t.2-.55l7.55-5.3q.525-.35 1.15-.35t1.15.35l7.55 5.3q.3.2.2.55t-.475.35H13v12q0 .425-.288.713T12 22Zm0-15h3.65h-7.3H12ZM4 22q-.425 0-.713-.288T3 21v-4.25l-.625-3.425q-.075-.425.15-.75t.65-.4q.4-.075.738.163t.412.637L4.9 16H8q.425 0 .713.288T9 17v4q0 .425-.288.713T8 22q-.425 0-.713-.288T7 21v-3H5v3q0 .425-.288.713T4 22Zm12 0q-.425 0-.713-.288T15 21v-4q0-.425.288-.713T16 16h3.1l.575-3.025q.075-.4.413-.638t.737-.162q.425.075.65.4t.15.75L21 16.75V21q0 .425-.288.713T20 22q-.425 0-.713-.288T19 21v-3h-2v3q0 .425-.288.713T16 22ZM8.35 7h7.3L12 4.45L8.35 7Z"/>`),
		g.Group(children),
	)
}