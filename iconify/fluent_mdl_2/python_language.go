package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PythonLanguage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M480 384q93 0 171 22t137 69t90 119t33 172q0 102-37 179t-100 129t-147 79t-180 27H278v484H128V384h352zm-45 660q67 0 125-14t101-47t68-84t25-126q0-71-22-119t-63-78t-95-43t-120-13H278v524h157zm1549-660l-422 807v473h-151v-470l-411-810h171l286 578q9 19 16 39t18 40q6-20 15-40t19-39l300-578h159z"/>`),
		g.Group(children),
	)
}