package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CovidcheckAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.642 7.252h8.506v8.506m0 17.012v8.505h-8.506m-17.012 0H7.124V32.77m10.936-8.506l4.496 4.86l9.478-9.721M5.588 5.308c2.706-1.371 9.448-1.383 12.152 0c.011 4.266-.007 9.925-6.076 12.758c-6.05-2.848-6.06-8.49-6.075-12.758Zm6.076 1.58v7.29m-3.645-3.767h7.29"/>`),
		g.Group(children),
	)
}