package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SocialForrst(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M84.173 91L49.968 9L15.827 90.853h28.801V73.984l-8.655-5.935l2.119-3.139l6.536 4.464V56.299h8.947v8.306l6.466-3.414l1.864 3.426l-8.33 4.407v4.708l12.537-6.274l1.665 3.377l-14.202 7.123v12.895h17.823V91z"/>`),
		g.Group(children),
	)
}