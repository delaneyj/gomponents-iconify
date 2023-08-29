package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M386 148q40 39 40 92t-40 90l-21-22q29-30 29-70t-29-68zm-45 45q20 21 20 47t-20 45l-21-22q18-24 0-49zM256 5q18 0 30.5 12.5T299 48v384q0 18-12.5 30.5T256 475H43q-18 0-30.5-12.5T0 432V48q0-18 12.5-30.5T43 5h213zm0 406V69H43v342h213z"/>`),
		g.Group(children),
	)
}