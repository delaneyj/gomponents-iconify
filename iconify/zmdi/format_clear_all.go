package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatClearAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 213v-42h298v42H43zM0 299v-43h299v43H0zM85 85h299v43H85V85z"/>`),
		g.Group(children),
	)
}