package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Movie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0v512h512V0H0zm63.9 490.7H21.3V448H64v42.7zm0-85.3H21.3v-42.7H64v42.7zm0-85.4H21.3v-42.7H64V320zm0-85.3H21.3V192H64v42.7zm0-85.4H21.3v-42.7H64v42.7zm0-85.4H21.3V21.3H64v42.6zm362.8 405.5H85.3V277.3h341.4v192.1zm0-234.7H85.3V42.6h341.4v192.1zm64 256H448V448h42.7v42.7zm0-85.3H448v-42.7h42.7v42.7zm0-85.4H448v-42.7h42.7V320zm0-85.3H448V192h42.7v42.7zm0-85.4H448v-42.7h42.7v42.7zm0-85.4H448V21.3h42.7v42.6z"/>`),
		g.Group(children),
	)
}