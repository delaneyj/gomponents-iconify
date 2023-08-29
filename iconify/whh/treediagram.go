package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Treediagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1024h-64q-27 0-45.5-18.5T768 960v-64q0-26 18.5-45t45.5-19V576H512v256q26 0 45 19t19 45v64q0 27-19 45.5t-45 18.5h-64q-27 0-45.5-18.5T384 960v-64q0-26 18.5-45t45.5-19V576H128v256q26 0 45 19t19 45v64q0 27-19 45.5t-45 18.5H64q-27 0-45.5-18.5T0 960v-64q0-26 18.5-45T64 832V544q0-13 9.5-22.5T96 512h352V320h-64q-27 0-45.5-18.5T320 256V64q0-26 18.5-45T384 0h192q26 0 45 19t19 45v192q0 27-19 45.5T576 320h-64v192h352q13 0 22.5 9.5T896 544v288q26 0 45 19t19 45v64q0 27-19 45.5t-45 18.5z"/>`),
		g.Group(children),
	)
}