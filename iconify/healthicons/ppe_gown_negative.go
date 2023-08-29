package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeGownNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPpeGownNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM35.34 8.526A5 5 0 0 0 30.363 4H27a3 3 0 1 1-6 0h-3.363a5 5 0 0 0-4.977 4.526l-1.447 15.19A2.988 2.988 0 0 0 12 26.04V30a1 1 0 0 0 1 1h3v12a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V31h3a1 1 0 0 0 1-1v-3.96c.56-.603.872-1.433.787-2.324L35.34 8.526ZM32 29h2v-2h-2v2Zm-16-2h-2v2h2v-2Zm15-2H17v-2h14v2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeGownNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}