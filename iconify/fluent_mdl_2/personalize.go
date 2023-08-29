package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Personalize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1792v128h256v128H640v-128h256v-128H0V640h1024L895 768H128v896h1664V896l128-128v1024h-896zm-512-384q27 0 50-10t40-27t28-41t10-50q0-49 17-93t48-78t72-56t91-28l711-711q28-28 65-43t76-15q41 0 77 15t64 43t42 63t16 79q0 39-15 76t-43 65l-711 711q-5 48-27 90t-56 72t-78 48t-93 18H256v-128h256zM1720 384q-29 0-50 21l-445 445q57 44 101 101l445-445q21-21 21-50q0-30-21-51t-51-21zm-678 649q61 40 101 101l92-91q-42-60-102-102l-91 92zm-18 247q0-27-10-50t-27-40t-41-28t-50-10q-27 0-50 10t-40 27t-28 41t-10 50q0 68-34 128h162q26 0 49-10t41-27t28-41t10-50z"/>`),
		g.Group(children),
	)
}