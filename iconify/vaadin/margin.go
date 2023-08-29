package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Margin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0h1v1H0V0zm2 0h1v1H2V0zM1 1h1v1H1V1zM0 2h1v1H0V2zm2 0h1v1H2V2zM1 3h1v1H1V3zM0 4h1v1H0V4zm1 1h1v1H1V5zM0 6h1v1H0V6zm1 1h1v1H1V7zM0 8h1v1H0V8zm1 1h1v1H1V9zm-1 1h1v1H0v-1zm1 1h1v1H1v-1zm-1 1h1v1H0v-1zm1 1h1v1H1v-1zm-1 1h1v1H0v-1zm2 0h1v1H2v-1zm-1 1h1v1H1v-1zm2 0h1v1H3v-1zm2 0h1v1H5v-1zM4 0h1v1H4V0zM3 1h1v1H3V1zm2 0h1v1H5V1zM4 14h1v1H4v-1zM6 0h1v1H6V0zm2 0h1v1H8V0zM7 1h1v1H7V1zM6 14h1v1H6v-1zm2 0h1v1H8v-1zm-1 1h1v1H7v-1zm2 0h1v1H9v-1zm2 0h1v1h-1v-1zM10 0h1v1h-1V0zM9 1h1v1H9V1zm2 0h1v1h-1V1zm-1 13h1v1h-1v-1zm2-14h1v1h-1V0zm2 0h1v1h-1V0zm-1 1h1v1h-1V1z"/><path fill="currentColor" d="M13 2h-1v1h-1V2h-1v1H9V2H8v1H7V2H6v1H5V2H4v1H3v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v-1h1v-1h-1v-1h1V9h-1V8h1V7h-1V6h1V5h-1V4h1V3h-1V2zm-1 10H4V4h8v8zm2-10h1v1h-1V2zm0 2h1v1h-1V4zm0 2h1v1h-1V6zm0 2h1v1h-1V8zm0 2h1v1h-1v-1zm0 2h1v1h-1v-1z"/><path fill="currentColor" d="M13 13h1v1h-1v-1zm-1 1h1v1h-1v-1zm2 0h1v1h-1v-1zm-1 1h1v1h-1v-1zm2 0h1v1h-1v-1zm0-14h1v1h-1V1zm0 2h1v1h-1V3zm0 2h1v1h-1V5zm0 2h1v1h-1V7zm0 2h1v1h-1V9zm0 2h1v1h-1v-1zm0 2h1v1h-1v-1z"/>`),
		g.Group(children),
	)
}