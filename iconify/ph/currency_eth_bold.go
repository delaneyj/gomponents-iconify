package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyEthBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m225.44 120.59l-88-112a12 12 0 0 0-18.88 0l-88 112a12 12 0 0 0 0 14.82l.6.76a3.72 3.72 0 0 0 .44.56l87 110.68a12 12 0 0 0 18.88 0l88-112a12 12 0 0 0-.04-14.82ZM140 50.7l57.12 72.7l-57.12 26Zm-24 98.66l-57.12-26L116 50.7Zm0 26.37v29.57l-36.15-46Zm24 0l36.15-16.43l-36.15 46Z"/>`),
		g.Group(children),
	)
}