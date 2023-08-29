package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SysProvision(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 20v10l9-5l-9-5z"/><circle cx="14" cy="14" r="2" fill="currentColor"/><path fill="currentColor" d="M14 20a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6Zm0-10a4 4 0 1 0 4 4a4.005 4.005 0 0 0-4-4Z"/><path fill="currentColor" d="M25.951 12.91h-.006A12.05 12.05 0 1 0 17 25.605v-2.066a9.981 9.981 0 1 1 6.623-6.81l1.925.544A12.034 12.034 0 0 0 26 14q0-.55-.049-1.09Z"/>`),
		g.Group(children),
	)
}