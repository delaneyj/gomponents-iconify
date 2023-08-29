package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1376 32q44 52 12 148t-108 172l-161 161l160 696q5 19-12 33l-128 96q-7 6-19 6q-4 0-7-1q-15-3-21-16L813 819l-259 259l53 194q5 17-8 31l-96 96q-9 9-23 9h-2q-15-2-24-13l-189-252L13 954q-11-7-13-23q-1-13 9-25l96-97q9-9 23-9q6 0 8 1l194 53l259-259L81 316q-14-8-17-24q-2-16 9-27l128-128q14-13 30-8l665 159l160-160q76-76 172-108t148 12z"/>`),
		g.Group(children),
	)
}