package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminalPowershell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m2.5 1.007l12.999.17l.43.501l-1.82 12.872l-.57.489l-13-.17l-.43-.502L1.93 1.495l.57-.488zM1.18 13.885l11.998.157l1.68-11.882L2.86 2.003L1.18 13.885zm5.791-3.49l-.14.991l5 .066l.14-.99l-5-.066zm1.71-2.457l-3.663-2.93l-.692.796l2.636 2.112L3.739 9.95l.465.812L8.68 7.938z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}