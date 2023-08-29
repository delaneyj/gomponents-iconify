package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SentimentVeryDissatisfiedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13.5q-1.7 0-3.088.963T6.9 17h10.2q-.625-1.575-2.013-2.538T12 13.5ZM7.8 12l1.1-1.05L9.95 12L11 10.95L9.95 9.9L11 8.8L9.95 7.75L8.9 8.8L7.8 7.75L6.75 8.8L7.8 9.9l-1.05 1.05L7.8 12Zm6.25 0l1.05-1.05L16.2 12l1.05-1.05L16.2 9.9l1.05-1.1l-1.05-1.05l-1.1 1.05l-1.05-1.05L13 8.8l1.05 1.1L13 10.95L14.05 12ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-10Zm0 8q3.35 0 5.675-2.325T20 12q0-3.35-2.325-5.675T12 4Q8.65 4 6.325 6.325T4 12q0 3.35 2.325 5.675T12 20Z"/>`),
		g.Group(children),
	)
}