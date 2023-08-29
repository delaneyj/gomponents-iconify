package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 1.932l9 3.375V10c0 4.118-1.62 7.113-3.535 9.074a11.407 11.407 0 0 1-2.915 2.175c-.913.468-1.826.751-2.55.751c-.724 0-1.637-.283-2.55-.75a11.41 11.41 0 0 1-2.915-2.176C4.619 17.113 3 14.118 3 10V5.307l9-3.375zm3.707 8.775a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l4-4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}