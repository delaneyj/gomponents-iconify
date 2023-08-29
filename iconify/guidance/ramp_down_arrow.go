package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RampDownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1 23h22v-1C9 18.772 1.5 13 1.5 13H1v10Zm22.5-8.746L11 7m12.5 7.254c-.498-.29-.944-.95-1.274-1.559a7.972 7.972 0 0 1-.856-2.622c-.106-.712-.155-1.514.068-1.902m2.062 6.083c-.498-.29-1.29-.348-1.98-.33a7.869 7.869 0 0 0-2.688.566c-.666.263-1.38.622-1.604 1.01"/>`),
		g.Group(children),
	)
}