package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParkOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.8 0-1.375-.575t-.575-1.375V18H4.875q-.6 0-.9-.525t.05-1.025L7 12h-.075q-.6 0-.888-.537t.063-1.038l5.075-7.25q.15-.2.375-.312T12 2.75q.225 0 .45.113t.375.312l5.075 7.25q.35.5.063 1.038t-.888.537H17l2.975 4.45q.35.5.05 1.025t-.9.525H13.95v2.05q0 .8-.575 1.375T12 22Zm-5.25-6h4h-1.9h6.3h-1.9h4h-10.5Zm0 0h10.5l-4-6h1.9L12 5.5L8.85 10h1.9l-4 6Z"/>`),
		g.Group(children),
	)
}