package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7ZM8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6Z"/><path fill="currentColor" d="M8 4.5c-1.11 0-2 .89-2 2h1c0-.55.45-1 1-1s1 .45 1 1c0 1-1.5.88-1.5 2.5h1c0-1.12 1.5-1.25 1.5-2.5c0-1.11-.89-2-2-2Z"/><circle cx="8" cy="11" r=".62" fill="currentColor"/><circle cx="6.5" cy="6.5" r=".5" fill="currentColor"/><circle cx="8" cy="9" r=".5" fill="currentColor"/>`),
		g.Group(children),
	)
}