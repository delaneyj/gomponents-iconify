package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M325.8 279.3H186.2v139.6h139.6V279.3zm46.6-139.7v279.3H512V139.6H372.4zM139.6 0H0v418.9h139.6V0zM0 512h512v-46.5H0V512z"/>`),
		g.Group(children),
	)
}