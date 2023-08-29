package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Firebase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5.188 20.896L8.339.615C8.443-.073 9.37-.229 9.693.386l3.391 6.359zm22.39 4.922l-3-18.667a.725.725 0 0 0-1.224-.391L4.422 25.817l10.474 5.906c.656.37 1.458.37 2.115 0zm-8.51-16.287l-2.427-4.646a.724.724 0 0 0-1.281 0L4.709 23.979z"/>`),
		g.Group(children),
	)
}