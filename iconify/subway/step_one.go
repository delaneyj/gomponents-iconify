package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M140.6 280.3H1v139.6h139.6V280.3zm186.2-139.7H187.2v279.3h139.6V140.6zM373.4 1v418.9H513V1H373.4zM1 513h512v-46.5H1V513z"/>`),
		g.Group(children),
	)
}