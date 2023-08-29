package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeerAltFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<mask id="lineMdBeerAltFilled0"><g transform="translate(0 14)"><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7.67C6.6 7.3 7.22 7 8 7C10 7 11 9 13 9C14.64 9 15.6 7.66 17 7.17"/><path fill="#fff" d="M17 8L16 21H7L6 8z"/><animateMotion fill="freeze" begin="0.6s" calcMode="linear" dur="0.6s" path="M0 0v-14"/></g></mask><path fill="none" stroke="currentColor" stroke-dasharray="60" stroke-dashoffset="60" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3L16 21H7L5 3z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="60;0"/></path><rect width="11" height="14" x="6" y="6" fill="currentColor" mask="url(#lineMdBeerAltFilled0)"/>`),
		g.Group(children),
	)
}