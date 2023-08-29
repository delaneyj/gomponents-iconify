package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerSimpleSlashThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188 152v-48a4 4 0 0 1 8 0v48a4 4 0 0 1-8 0Zm36-68a4 4 0 0 0-4 4v80a4 4 0 0 0 8 0V88a4 4 0 0 0-4-4Zm-13 129.31a4 4 0 1 1-5.92 5.38l-49-53.94V224a4 4 0 0 1-2.24 3.59a3.92 3.92 0 0 1-1.84.41a4.07 4.07 0 0 1-2.46-.84L78.63 172H32a12 12 0 0 1-12-12V96a12 12 0 0 1 12-12h50.59L45 42.69a4 4 0 0 1 6-5.38ZM148 156L89.87 92H32a4 4 0 0 0-4 4v64a4 4 0 0 0 4 4h48a4.06 4.06 0 0 1 2.46.84l65.54 51Zm-33.39-89.8l33.39-26v66.65a4 4 0 0 0 8 0V32a4 4 0 0 0-6.46-3.16l-39.85 31a4 4 0 1 0 4.92 6.31Z"/>`),
		g.Group(children),
	)
}