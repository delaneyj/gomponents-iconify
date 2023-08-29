package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeFundsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.379 15.106A8.001 8.001 0 0 0 8.035 5.056l-.993-1.737a9.996 9.996 0 0 1 9.962.023c4.49 2.593 6.21 8.143 4.118 12.77l1.341.775l-4.165 2.214l-.165-4.714l1.246.719ZM4.629 8.9a8.001 8.001 0 0 0 11.345 10.05l.992 1.737a9.995 9.995 0 0 1-9.962-.024C2.514 18.07.794 12.52 2.887 7.893L1.545 7.12L5.71 4.905l.165 4.714l-1.245-.72Zm8.79 5.931l-2.829-2.828l-2.828 2.828l-1.415-1.414l4.243-4.243l2.828 2.829l2.829-2.829l1.414 1.414l-4.243 4.243Z"/>`),
		g.Group(children),
	)
}