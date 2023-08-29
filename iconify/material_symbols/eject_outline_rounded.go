package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19q-.425 0-.713-.288T5 18q0-.425.288-.713T6 17h12q.425 0 .713.288T19 18q0 .425-.288.713T18 19H6Zm1.225-4q-.6 0-.888-.525t.038-1.025l4.8-7.2q.3-.45.825-.45t.825.45l4.8 7.2q.325.5.038 1.025t-.888.525h-9.55ZM12 13Zm-2.95 0h5.9L12 8.6L9.05 13Z"/>`),
		g.Group(children),
	)
}