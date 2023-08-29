package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peso(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11.5 7h-6c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h6c.28 0 .5.22.5.5s-.22.5-.5.5Zm0 2h-6c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h6c.28 0 .5.22.5.5s-.22.5-.5.5Z"/><path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7ZM8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6Z"/><path fill="currentColor" d="M6.5 13c-.28 0-.5-.22-.5-.5v-7C6 4.67 6.67 4 7.5 4h1A2.5 2.5 0 0 1 11 6.5v2A2.5 2.5 0 0 1 8.5 11H7v1.5c0 .28-.22.5-.5.5Zm.5-3h1.5c.83 0 1.5-.67 1.5-1.5v-2C10 5.67 9.33 5 8.5 5h-1c-.28 0-.5.22-.5.5V10Z"/>`),
		g.Group(children),
	)
}