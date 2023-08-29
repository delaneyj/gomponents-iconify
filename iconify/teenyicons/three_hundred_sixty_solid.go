package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeHundredSixtySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 8.5V8h.5a.5.5 0 1 1-.5.5ZM11.5 6a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 1 0v-2a.5.5 0 0 0-.5-.5Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 0 1 12.787-5.32a1.5 1.5 0 0 1 1.659 2.484A7.52 7.52 0 0 1 15 7.5a7.5 7.5 0 0 1-15 0ZM13.5 3a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1ZM2 6h1.5l-.9 1.2A.5.5 0 0 0 3 8h.5a.5.5 0 0 1 0 1H2v1h1.5a1.5 1.5 0 0 0 .449-2.932L4.9 5.8a.5.5 0 0 0-.4-.8H2v1Zm5.5-1A1.5 1.5 0 0 0 6 6.5v2A1.5 1.5 0 1 0 7.5 7H7v-.5a.5.5 0 0 1 .5-.5H8V5h-.5ZM10 6.5a1.5 1.5 0 0 1 3 0v2a1.5 1.5 0 0 1-3 0v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}