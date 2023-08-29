package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotationNoneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.85 20.75q-.3-.3-.3-.713t.3-.712L17.2 19H4.975q-.425 0-.7-.288T4 18q0-.425.288-.713T5 17h12.2l-.375-.35q-.275-.275-.262-.7t.287-.7q.275-.275.7-.275t.7.275l2.025 2.05q.15.15.225.338t.075.387q0 .2-.062.388t-.213.337l-2.025 2.025q-.3.3-.712.288t-.713-.313ZM8.175 14q-.55 0-.788-.363t-.037-.862l3.3-8.85q.15-.375.537-.65T12 3q.425 0 .8.275t.525.65l3.3 8.85q.2.5-.05.863t-.8.362q-.25 0-.475-.163t-.3-.387l-.75-2.25H9.8l-.825 2.25q-.075.225-.3.388t-.5.162Zm2.175-4.4h3.3l-1.6-4.55h-.1l-1.6 4.55Z"/>`),
		g.Group(children),
	)
}