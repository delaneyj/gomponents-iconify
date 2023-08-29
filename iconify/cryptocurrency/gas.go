package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm9-9.42V9.621l-6.99 2.48v7.22L25 22.58zM14.823 26V13.139L8 9.958V22.82L14.823 26zm10.01-16.843l.061-.021L18.165 6l-.062.021l-.177.063l-.084.03L8.27 9.51l6.729 3.136l2.843-1.008l.167-.06l6.711-2.38l.112-.04z"/>`),
		g.Group(children),
	)
}