package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceHierarchyThreeNodeOrganizationLinksStructureLinkNodesNetworkHierarchy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="3.5" height="3.5" x=".5" y="10" rx=".5"/><rect width="3.5" height="3.5" x="10" y="10" rx=".5"/><rect width="4" height="4" x="5" y=".5" rx=".5"/><path d="M4 12h6M5.09 4.29L2.5 10m6.41-5.71L11.5 10"/></g>`),
		g.Group(children),
	)
}