package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceHierarchyOneNodeOrganizationLinksStructureLinkNodesNetworkHierarchy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="2.5" cy="7" r="2"/><circle cx="11.5" cy="2.5" r="2"/><circle cx="11.5" cy="11.5" r="2"/><path d="m3.71 5.41l5.85-2.43M3.71 8.59l5.85 2.43"/></g>`),
		g.Group(children),
	)
}