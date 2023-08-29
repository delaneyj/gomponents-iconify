package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1200 100.232V276.38H0V100.306l1200-.074zM0 374.744l337.939 225.293L0 825.33V374.744zm1200 1.318v175.561H506.25V376.062H1200zm0 273.706V825.33H506.25V649.768H1200zM0 923.693l1200 .073v176.001H0V923.693z"/>`),
		g.Group(children),
	)
}