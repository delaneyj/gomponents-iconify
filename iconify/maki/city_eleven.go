package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.751 3h-.75V1.25a.25.25 0 0 0-.25-.25h-.5a.25.25 0 0 0-.25.25V3h-.75a.25.25 0 0 0-.25.25v3.751h-1.75a.251.251 0 0 0-.25.251v2.5a.251.251 0 0 0 .25.249h4.501a.25.25 0 0 0 .25-.25V3.25a.25.25 0 0 0-.25-.25zm-3.75 5.001h1v1H6v-1zm3 1H8v-1h1v1zm0-2H8v-1h1v1zm0-2H8v-1h1v1zm-3-3.75A.25.25 0 0 0 5.75 1H3.25a.25.25 0 0 0-.25.25V2h-.75a.25.25 0 0 0-.25.25V3h-.75a.25.25 0 0 0-.25.25v6.501c0 .138.112.25.25.25H4v-4h2V1.25zM3 9H2v-1h1v1zm0-2H2v-1h1v1zm0-2H2v-1h1v1zm2 0H4v-1h1v1zm0-2H4V2h1v1.001z" fill="currentColor"/>`),
		g.Group(children),
	)
}