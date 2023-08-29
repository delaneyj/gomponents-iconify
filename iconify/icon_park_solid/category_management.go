package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CategoryManagement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCategoryManagement0"><g fill="none"><rect width="36" height="14" x="6" y="28" stroke="#fff" stroke-width="4" rx="4"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M20 7H10a4 4 0 0 0-4 4v6a4 4 0 0 0 4 4h10"/><circle cx="34" cy="14" r="8" fill="#fff" stroke="#fff" stroke-width="4"/><circle cx="34" cy="14" r="3" fill="#000"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCategoryManagement0)"/>`),
		g.Group(children),
	)
}