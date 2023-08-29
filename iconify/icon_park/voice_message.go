package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceMessage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" d="M24 44C35.0457 44 44 35.0457 44 24C44 12.9543 35.0457 4 24 4C12.9543 4 4 12.9543 4 24C4 35.0457 12.9543 44 24 44Z"/><path fill="#fff" d="M17 25.8994C18.1046 25.8994 19 25.004 19 23.8994C19 22.7948 18.1046 21.8994 17 21.8994C15.8954 21.8994 15 22.7948 15 23.8994C15 25.004 15.8954 25.8994 17 25.8994Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21.9497 28.8492C23.2165 27.5825 24 25.8325 24 23.8995C24 21.9665 23.2165 20.2165 21.9497 18.9497"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M26.8994 33.799C29.4329 31.2655 30.9999 27.7655 30.9999 23.8995C30.9999 20.0335 29.4329 16.5335 26.8994 14"/></g>`),
		g.Group(children),
	)
}