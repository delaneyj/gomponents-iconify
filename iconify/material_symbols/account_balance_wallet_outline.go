package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountBalanceWalletOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13.5q.65 0 1.075-.425T17.5 12q0-.65-.425-1.075T16 10.5q-.65 0-1.075.425T14.5 12q0 .65.425 1.075T16 13.5ZM5 19V5v14Zm0 2q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v2.5h-2V5H5v14h14v-2.5h2V19q0 .825-.588 1.413T19 21H5Zm8-4q-.825 0-1.413-.588T11 15V9q0-.825.588-1.413T13 7h7q.825 0 1.413.588T22 9v6q0 .825-.588 1.413T20 17h-7Zm7-2V9h-7v6h7Z"/>`),
		g.Group(children),
	)
}