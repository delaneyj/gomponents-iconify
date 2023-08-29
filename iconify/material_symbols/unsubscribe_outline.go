package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnsubscribeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13L4 8v10h8q0 .525.075 1.012T12.3 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v5.7q-.45-.225-.963-.375T20 11.1V8l-8 5Zm0-2l8-5H4l8 5Zm7 12q-2.075 0-3.538-1.463T14 18q0-2.075 1.463-3.538T19 13q2.075 0 3.538 1.463T24 18q0 2.075-1.463 3.538T19 23Zm-3-4.5h6v-1h-6v1ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}