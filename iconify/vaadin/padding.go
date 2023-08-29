package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Padding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm15 3h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v-1h1v-1H1v-1h1v-1H1V9h1V8H1V7h1V6H1V5h1V4H1V3h1V2H1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1v1z"/><path fill="currentColor" d="M3 2h1v1H3V2zm1 1h1v1H4V3zm2 0h1v1H6V3zM5 2h1v1H5V2zm2 0h1v1H7V2zm2 0h1v1H9V2zM8 3h1v1H8V3zm2 0h1v1h-1V3zm2 0h1v1h-1V3zm-1-1h1v1h-1V2zm2 0h1v1h-1V2zm-1 3h1v1h-1V5zm1-1h1v1h-1V4zm-1 3h1v1h-1V7zm1-1h1v1h-1V6zm-1 3h1v1h-1V9zm1-1h1v1h-1V8zm-1 3h1v1h-1v-1zm1-1h1v1h-1v-1zm-1 3h1v1h-1v-1zm1-1h1v1h-1v-1zM2 3h1v1H2V3zm1 1h1v1H3V4zM2 5h1v1H2V5zm1 1h1v1H3V6zM2 7h1v1H2V7zm1 1h1v1H3V8zM2 9h1v1H2V9zm1 1h1v1H3v-1zm-1 1h1v1H2v-1zm0 2h1v1H2v-1zm1-1h1v1H3v-1zm1-1h1v1H4v-1zm0 2h1v1H4v-1zm1-1h1v1H5v-1zm1 1h1v1H6v-1zm1-1h1v1H7v-1zm2 0h1v1H9v-1zm-1 1h1v1H8v-1zm3-1h1v1h-1v-1zm-1 1h1v1h-1v-1z"/>`),
		g.Group(children),
	)
}