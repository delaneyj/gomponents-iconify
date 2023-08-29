package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7ZM8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6Z"/><path fill="currentColor" d="M6.5 13c-.28 0-.5-.22-.5-.5V9H4.5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H6V4.5c0-.28.22-.5.5-.5h3a2.5 2.5 0 0 1 0 5H7v3.5c0 .28-.22.5-.5.5ZM7 8h2.5c.83 0 1.5-.67 1.5-1.5S10.33 5 9.5 5H7v3Z"/><path fill="currentColor" d="M9.5 11h-5c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h5c.28 0 .5.22.5.5s-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}