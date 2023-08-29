package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackSquareButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#405866" d="M63.932 57.1a6.892 6.892 0 0 1-6.893 6.896H6.889A6.893 6.893 0 0 1 0 57.1V6.96A6.892 6.892 0 0 1 6.889.067h50.15a6.892 6.892 0 0 1 6.893 6.893V57.1"/><path fill="#fff" d="M46.48 13.524H17.44a3.99 3.99 0 0 0-3.99 3.99v29.04a3.994 3.994 0 0 0 3.99 3.994h29.04a3.993 3.993 0 0 0 3.989-3.994v-29.04a3.99 3.99 0 0 0-3.989-3.99"/>`),
		g.Group(children),
	)
}