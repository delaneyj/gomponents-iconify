package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 100.231v176.147h1200V100.305L0 100.231zm1200 274.512L862.061 600.036L1200 825.329V374.743zM0 376.062v175.562h693.75V376.062H0zm0 273.706V825.33h693.75V649.768H0zm1200 273.926L0 923.768v176.001h1200V923.694z"/>`),
		g.Group(children),
	)
}