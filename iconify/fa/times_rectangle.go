package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimesRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1175 1193l146-146q10-10 10-23t-10-23l-233-233l233-233q10-10 10-23t-10-23l-146-146q-10-10-23-10t-23 10L896 576L663 343q-10-10-23-10t-23 10L471 489q-10 10-10 23t10 23l233 233l-233 233q-10 10-10 23t10 23l146 146q10 10 23 10t23-10l233-233l233 233q10 10 23 10t23-10zm617-1033v1216q0 66-47 113t-113 47H160q-66 0-113-47T0 1376V160Q0 94 47 47T160 0h1472q66 0 113 47t47 113z"/>`),
		g.Group(children),
	)
}