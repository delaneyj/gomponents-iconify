package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoldMedal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M44 4H31L27 14.3001C31.4614 15.2057 35.2038 18.0914 37.2699 22L44 4Z"/><path fill="#2F88FF" stroke="#000" d="M17 4H4L10.7301 22C12.7962 18.0914 16.5386 15.2057 21 14.3001L17 4Z"/><path fill="#2F88FF" stroke="#000" d="M39 29C39 37.2843 32.2843 44 24 44C15.7157 44 9 37.2843 9 29C9 26.4718 9.62546 24.0897 10.7301 22C12.7962 18.0914 16.5386 15.2057 21 14.3001C21.9693 14.1033 22.9726 14 24 14C25.0274 14 26.0307 14.1033 27 14.3001C31.4614 15.2057 35.2038 18.0914 37.2699 22C38.3745 24.0897 39 26.4718 39 29Z"/><path stroke="#fff" d="M24 35V22L21 23M24 35H28M24 35H20"/></g>`),
		g.Group(children),
	)
}