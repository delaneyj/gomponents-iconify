package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoEnhanceOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1984q0-26 19-45L1235 723q19-19 45-19t45 19t19 45q0 26-19 45L109 2029q-19 19-45 19t-45-19t-19-45zM1408 0h128v256h-128V0zm-207 395l-182-181l91-91l181 182l-90 90zm-49 245H896V512h256v128zm256 256h128v256h-128V896zm335-139l182 181l-91 91l-181-182l90-90zm305-245v128h-256V512h256zm-305-117l-90-90l181-182l91 91l-182 181zm-271 117q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19z"/>`),
		g.Group(children),
	)
}