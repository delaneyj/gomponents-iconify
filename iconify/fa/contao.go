package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contao(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M138 0h197q-70 64-126 149q-36 56-59 115t-30 125.5t-8.5 120t10.5 132t21 126T171 904q4 19 6 28q51 238 81 329q57 171 152 275H138q-48 0-82-34t-34-82V116q0-48 34-82t82-34zm1208 0h308q48 0 82 34t34 82v1304q0 48-34 82t-82 34h-178q212-210 196-565l-469 101q-2 45-12 82t-31 72t-59.5 59.5t-93.5 36.5q-123 26-199-40q-32-27-53-61t-51.5-129T639 834q-35-163-45.5-263T588 432t23-77q20-41 62.5-73T776 237q45-12 83.5-6.5t67 17t54 35t43 48T1058 387l468-100q-68-175-180-287z"/>`),
		g.Group(children),
	)
}