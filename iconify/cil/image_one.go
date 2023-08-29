package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M1.875 1.875v20.25h20.25V1.875zm18.75 1.5v6.454l-3.036-3.036l-3.245 3.245L9.656 5.35l-6.281 6.281V3.374zm0 12.944l-5.221-5.221l2.185-2.185l3.036 3.036zm-17.25 4.306v-6.872l6.281-6.281l10.969 10.969v2.185z" fill="currentColor"/>`),
		g.Group(children),
	)
}