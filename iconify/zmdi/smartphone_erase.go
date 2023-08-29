package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneErase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m213 159l-85 85l85 86l-21 21l-85-85l-86 85l-21-21l85-86l-85-85l21-21l86 85l85-85zM341 5q18 0 30.5 12.5T384 48v384q0 18-12.5 30.5T341 475H128q-18 0-30.5-12.5T85 432v-64h43v43h213V69H128v43H85V48q0-18 12.5-30.5T128 5h213z"/>`),
		g.Group(children),
	)
}