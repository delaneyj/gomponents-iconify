package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatValignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m85 389l86-85l85 85h-64v86h-43v-86H85zM256 91l-85 85l-86-85h64V5h43v86h64zM0 219h341v42H0v-42z"/>`),
		g.Group(children),
	)
}