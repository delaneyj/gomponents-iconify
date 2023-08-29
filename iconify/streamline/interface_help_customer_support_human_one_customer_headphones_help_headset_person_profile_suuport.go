package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceHelpCustomerSupportHumanOneCustomerHeadphonesHelpHeadsetPersonProfileSuuport(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="7" r="2.5"/><path d="M7 4.5v-4m2.5 13V12h2a1 1 0 0 0 1-1V7a5.5 5.5 0 0 0-3-4.9m-5 0a5.5 5.5 0 0 0 0 9.8v1.6"/></g>`),
		g.Group(children),
	)
}