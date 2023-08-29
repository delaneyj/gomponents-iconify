package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeerAltTwotoneLoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdBeerAltTwotoneLoop0"><g transform="translate(0 14)"><path stroke="#fff" stroke-width="2" d="M18 7C16 7 15 9 13 9C11 9 10 7 8 7C6 7 5 9 3 9C1 9 0 7 -2 7C-4 7 -5 9 -7 9"><animateMotion calcMode="linear" dur="3s" path="M0 0h10" repeatCount="indefinite"/></path><path fill="#fff" fill-opacity=".3" d="M17 8L16 21H7L6 8z"/><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-14"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><rect width="11" height="14" x="6" y="6" fill="currentColor" mask="url(#lineMdBeerAltTwotoneLoop0)"/>`),
		g.Group(children),
	)
}