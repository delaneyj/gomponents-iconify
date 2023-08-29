package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Downdetector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 45c11.598 0 21-9.402 21-21S35.598 3 24 3S3 12.402 3 24v21h21Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="23.993" cy="35.723" r="2.777"/><path d="M20.396 10.22c2.795-1.084 5.103-.827 7.208 0c0 6.683-.485 13.137-1.03 19.563c-1.776.684-3.477.513-5.148 0c-.472-6.477-.977-12.945-1.03-19.563Z"/></g>`),
		g.Group(children),
	)
}