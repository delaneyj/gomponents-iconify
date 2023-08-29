package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Escalators(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEscalators0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M35 13L13 35H4v9h9l22-22h9v-9h-9Z"/><path d="m19 13l9-9m-6 0h6v6m-14 8l-9 9m6 0H5v-6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEscalators0)"/>`),
		g.Group(children),
	)
}