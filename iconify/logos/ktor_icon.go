package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KtorIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosKtorIcon0" x1="23.965%" x2="74.641%" y1="23.965%" y2="74.641%"><stop offset="29.6%" stop-color="#00AFFF"/><stop offset="69.4%" stop-color="#5282FF"/><stop offset="100%" stop-color="#945DFF"/></linearGradient><linearGradient id="logosKtorIcon1" x1="26.206%" x2="74.117%" y1="26.206%" y2="74.117%"><stop offset="10.8%" stop-color="#C757BC"/><stop offset="17.3%" stop-color="#CD5CA9"/><stop offset="49.2%" stop-color="#E8744F"/><stop offset="71.6%" stop-color="#F88316"/><stop offset="82.3%" stop-color="#FF8900"/></linearGradient></defs><path fill="url(#logosKtorIcon0)" d="m170.667 85.333l-42.49-42.489L85.333 0L40.277 45.056L0 85.333l85.333 85.334z"/><path fill="url(#logosKtorIcon1)" d="m85.333 170.667l42.49 42.489L170.667 256l45.056-45.056L256 170.667l-85.333-85.334z"/><path d="M170.667 85.333H85.333v85.334h85.334z"/>`),
		g.Group(children),
	)
}