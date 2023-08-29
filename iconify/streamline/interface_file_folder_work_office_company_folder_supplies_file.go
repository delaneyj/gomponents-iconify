package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceFileFolderWorkOfficeCompanyFolderSuppliesFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m.91 12.56l-.41-7A.5.5 0 0 1 1 5h4.61a.51.51 0 0 1 .49.38L6.5 7H13a.5.5 0 0 1 .5.54l-.39 5a1 1 0 0 1-1 .92H1.91a1 1 0 0 1-1-.9ZM3.5 3V1A.5.5 0 0 1 4 .5h8.5a.5.5 0 0 1 .5.5v4M7.5 3h3"/>`),
		g.Group(children),
	)
}