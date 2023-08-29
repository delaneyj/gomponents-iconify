package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileHiding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M10 44H38C39.1046 44 40 43.1046 40 42V14H30V4H10C8.89543 4 8 4.89543 8 6V42C8 43.1046 8.89543 44 10 44Z"/><path stroke="#000" d="M30 4L40 14"/><path stroke="#fff" d="M16 23C16.2821 23.9145 16.7095 24.7628 17.2546 25.5166C18.7827 27.63 21.2351 29 24 29C26.7649 29 29.2173 27.63 30.7454 25.5166C31.2905 24.7628 31.7179 23.9145 32 23"/><path stroke="#fff" d="M21.5215 29.0684L20.4862 32.9321"/><path stroke="#fff" d="M26.4863 29.0684L27.5216 32.9321"/><path stroke="#fff" d="M30.3535 27.3536L33.1819 30.182"/><path stroke="#fff" d="M15 30.0103L17.8285 27.1819"/></g>`),
		g.Group(children),
	)
}