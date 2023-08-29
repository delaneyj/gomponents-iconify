package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiPhoneservice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.372 7.686a20.886 20.886 0 0 0 6.683 5.514l11.657-.293a42.385 42.385 0 0 1 1.523 25.424c-7.659 3.464-16.35 4.153-25.775 2.87c-5.208-5.04-7.676-11.162-9.57-17.867c-.587-2.082-.906-5.634 2.036-6.042c3.63-.504 11.385.284 11.207-.282c-4.293-2.93-2.602-14.973 2.239-9.325Zm9.085 5.454l1.816 18.807"/>`),
		g.Group(children),
	)
}