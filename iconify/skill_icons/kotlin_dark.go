package skill_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KotlinDark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="none"><rect width="256" height="256" fill="#242938" rx="60"/><path fill="url(#skillIconsKotlinDark0)" d="M218 218H38V38h180l-90 90l90 90Z"/><defs><linearGradient id="skillIconsKotlinDark0" x1="218" x2="38" y1="38" y2="218" gradientUnits="userSpaceOnUse"><stop offset=".003" stop-color="#E44857"/><stop offset=".469" stop-color="#C711E1"/><stop offset="1" stop-color="#7F52FF"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}