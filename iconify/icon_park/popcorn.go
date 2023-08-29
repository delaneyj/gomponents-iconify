package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Popcorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M33.6961 40.8677L39 17H7L12.3039 40.8677C12.6376 42.3694 12.8045 43.1202 13.3529 43.5601C13.9013 44 14.6704 44 16.2087 44H29.7913C31.3296 44 32.0987 44 32.6471 43.5601C33.1955 43.1202 33.3624 42.3694 33.6961 40.8677Z"/><path stroke="#fff" d="M27 44L28 17"/><path stroke="#fff" d="M19 44L18 17"/><path stroke="#000" d="M31 44H15"/><path stroke="#000" d="M31 17H15"/><path stroke="#000" d="M11.0004 16.9999C11.0004 16.9999 10.0004 13.9999 11.5004 12.4999C13.0004 10.9999 16.0004 11.4999 16.0004 11.4999C16.0004 11.4999 16.0004 8.4999 18.0004 7.4999C20.0004 6.4999 23.0004 7.99988 23.0004 7.99988C23.0004 7.99988 25.0004 4.64271 28.0004 5.49992C31.0004 6.35713 31.0004 9.99994 31.0004 9.99994C31.0004 9.99994 33.5004 9.99998 35.0004 11.9999C36.5004 13.9999 35.0004 16.9999 35.0004 16.9999"/></g>`),
		g.Group(children),
	)
}