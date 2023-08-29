package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubstackReader(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.025 10.082V5.5H7.747v4.582h32.278Zm0 8.268v-4.58H7.747v4.581h32.278Zm.231 24.15V22.318c0-.164-.973-.108-1.887-.108H9.402c-.914 0-1.655-.055-1.655.108v19.888c.281-.163 15.341-8.212 16.169-8.598c.516.297 13.646 7.338 16.34 8.891h0Z"/>`),
		g.Group(children),
	)
}