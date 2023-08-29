package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.25 4A4.25 4.25 0 0 0 12 8.25v31.5A4.25 4.25 0 0 0 16.25 44h15.5A4.25 4.25 0 0 0 36 39.75V8.25A4.25 4.25 0 0 0 31.75 4h-15.5Zm5 31.5h5.5a1.25 1.25 0 1 1 0 2.5h-5.5a1.25 1.25 0 1 1 0-2.5Z"/>`),
		g.Group(children),
	)
}