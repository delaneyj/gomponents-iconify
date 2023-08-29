package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacialCleanser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M15 34v10h18V34"/><path d="M31.935 4h-15.95C9.043 4 8.514 8.605 9.25 11.04L14.989 34h17.942s4.644-17.817 5.888-22.96c.589-2.435.09-7.057-6.884-7.04Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M20.488 19.09c1.64-1.944 2.903-4.631 3.586-6.09c1.196 1.459 3.792 5.118 4.612 7.063c1.025 2.432-1.537 5.35-4.612 5.35c-3.074 0-5.636-3.89-3.586-6.322Z"/></g>`),
		g.Group(children),
	)
}