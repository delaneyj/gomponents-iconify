package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceHierarchyTwoNodeOrganizationLinksStructureLinkNodesNetworkHierarchy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="7" cy="2.5" r="1.5"/><circle cx="2" cy="11.5" r="1.5"/><circle cx="7" cy="11.5" r="1.5"/><circle cx="12" cy="11.5" r="1.5"/><path d="M2 10V8a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v2M7 4v6"/></g>`),
		g.Group(children),
	)
}