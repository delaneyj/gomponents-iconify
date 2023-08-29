package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightCircleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2zm0 1.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17zM15.25 8a.75.75 0 0 1 .743.648L16 8.75v6.5a.75.75 0 0 1-.648.743L15.25 16h-6.5a.75.75 0 0 1-.102-1.493l.102-.007h4.74L8.223 9.283a.75.75 0 0 1 .971-1.138l.085.072l5.222 5.172V8.75a.75.75 0 0 1 .75-.75z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}