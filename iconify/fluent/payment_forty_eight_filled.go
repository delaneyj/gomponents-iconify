package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaymentFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.25 9A6.25 6.25 0 0 0 4 15.25v3.25h40v-3.25A6.25 6.25 0 0 0 37.75 9h-27.5ZM4 32.75V21h40v11.75A6.25 6.25 0 0 1 37.75 39h-27.5A6.25 6.25 0 0 1 4 32.75Zm27.25-3.25a1.25 1.25 0 1 0 0 2.5h5.5a1.25 1.25 0 1 0 0-2.5h-5.5Z"/>`),
		g.Group(children),
	)
}