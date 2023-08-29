package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFormatRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19q-.425 0-.713-.288T5 18q0-.425.288-.713T6 17h12q.425 0 .713.288T19 18q0 .425-.288.713T18 19H6Zm4.35-8.4h3.3l-1.6-4.55h-.1l-1.6 4.55ZM8.175 15q-.55 0-.788-.363t-.037-.862l3.3-8.85q.15-.375.537-.65T12 4q.425 0 .8.275t.525.65l3.3 8.85q.2.5-.05.863t-.8.362q-.25 0-.475-.163t-.3-.387l-.75-2.25H9.8l-.825 2.25q-.075.225-.3.388t-.5.162Z"/>`),
		g.Group(children),
	)
}