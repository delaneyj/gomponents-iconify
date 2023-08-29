package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ktweak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.15 21.98s-3.538 4.505-7.86 4.505a5.543 5.543 0 1 0 0 11.086c9.539 0 9.896-11.494 7.86-15.59Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.478 4.5S38.53 10.553 38.53 29.001a14.53 14.53 0 0 1-29.06 0c0-11.447 5.84-16.312 5.84-16.312c-.043 4.063 2.835 6.935 6.385 6.935a6.421 6.421 0 0 0 6.429-6.415c0-6.715-1.645-8.709-1.645-8.709"/>`),
		g.Group(children),
	)
}