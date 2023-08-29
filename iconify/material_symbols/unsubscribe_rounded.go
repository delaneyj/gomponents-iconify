package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnsubscribeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11L5.3 6.8q-.425-.275-.862-.025T4 7.525q0 .225.1.413t.3.312l7.075 4.425q.25.15.525.15t.525-.15L19.6 8.25q.2-.125.3-.313t.1-.412q0-.5-.437-.75T18.7 6.8L12 11Zm7 12q-2.075 0-3.538-1.463T14 18q0-2.075 1.463-3.538T19 13q2.075 0 3.538 1.463T24 18q0 2.075-1.463 3.538T19 23Zm-2.5-4.5h5q.2 0 .35-.15T22 18q0-.2-.15-.35t-.35-.15h-5q-.2 0-.35.15T16 18q0 .2.15.35t.35.15ZM4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v5.7q-.7-.35-1.45-.525T19 11q-2.925 0-4.963 2.037T12 18q0 .5.075 1t.225 1H4Z"/>`),
		g.Group(children),
	)
}