package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CottageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19v-7.375L3 12.4q-.35.25-.75.2t-.65-.4q-.25-.35-.2-.75t.375-.65L4 9.1V7q0-.425.288-.713T5 6q.425 0 .713.288T6 7v.575l4.775-3.65Q11.325 3.5 12 3.5t1.225.425l9 6.875q.325.25.375.65t-.2.75q-.25.325-.65.375t-.725-.2L20 11.625V19q0 .825-.587 1.413T18 21H6q-.825 0-1.413-.588T4 19Zm2 0h5v-3q0-.425.288-.713T12 15q.425 0 .713.288T13 16v3h5v-8.9l-6-4.575L6 10.1V19Zm0 0h12H6ZM5.3 5q-.575 0-.888-.475t-.037-.975q.425-.725 1.113-1.137T7 2q.275 0 .525-.138t.375-.387q.125-.225.338-.35T8.725 1q.575 0 .875.475t.025.975q-.425.725-1.112 1.137T7 4q-.275 0-.525.125t-.375.4q-.125.225-.325.35T5.3 5Z"/>`),
		g.Group(children),
	)
}