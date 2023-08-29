package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1664 1248V480q-32 36-69 66q-268 206-426 338q-51 43-83 67t-86.5 48.5T897 1024h-2q-48 0-102.5-24.5T706 951t-83-67Q465 752 197 546q-37-30-69-66v768q0 13 9.5 22.5t22.5 9.5h1472q13 0 22.5-9.5t9.5-22.5zm0-1051v-24.5l-.5-13l-3-12.5l-5.5-9l-9-7.5l-14-2.5H160q-13 0-22.5 9.5T128 160q0 168 147 284q193 152 401 317q6 5 35 29.5t46 37.5t44.5 31.5T852 887t43 9h2q20 0 43-9t50.5-27.5T1035 828t46-37.5t35-29.5q208-165 401-317q54-43 100.5-115.5T1664 197zm128-37v1088q0 66-47 113t-113 47H160q-66 0-113-47T0 1248V160Q0 94 47 47T160 0h1472q66 0 113 47t47 113z"/>`),
		g.Group(children),
	)
}