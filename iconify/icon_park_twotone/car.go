package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCar0"><g fill="none"><path fill="#fff" fill-rule="evenodd" d="M13.5 32a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm21 0a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M7 37a3 3 0 0 1-3-3v-9.29a6 6 0 0 1 3.319-5.368l.682-.34l2.31-9.91A4 4 0 0 1 14.205 6h19.688a4 4 0 0 1 3.904 3.128l2.205 9.873l.68.34A6 6 0 0 1 44 24.709V34a3 3 0 0 1-3 3h-1.997v1A4.001 4.001 0 0 1 31 38v-1H17v1a4 4 0 1 1-8 0v-1H7Z"/><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M14 22h20l-1.652-7.434A2 2 0 0 0 30.396 13H17.604a2 2 0 0 0-1.952 1.566L14 22Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCar0)"/>`),
		g.Group(children),
	)
}