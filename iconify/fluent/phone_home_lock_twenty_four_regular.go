package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneHomeLockTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M17.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zM13.75 2A2.25 2.25 0 0 1 16 4.25v6.924a6.455 6.455 0 0 0-1.5.558V4.25a.75.75 0 0 0-.75-.75h-7.5a.75.75 0 0 0-.75.75v15.5c0 .414.336.75.75.75l5.483.001c.287.551.651 1.056 1.077 1.5L6.25 22A2.25 2.25 0 0 1 4 19.75V4.25A2.25 2.25 0 0 1 6.25 2h7.5zm3.75 11.998a1.75 1.75 0 0 0-1.744 1.607l-.006.143v.782a1 1 0 0 0-.743.853l-.007.115v2.5a1 1 0 0 0 .883.993l.117.007h3a1 1 0 0 0 .993-.883l.007-.117v-2.5a1 1 0 0 0-.644-.935l-.105-.033l-.001-.782a1.75 1.75 0 0 0-1.75-1.75zm-8.751 3.506L11 17.499c0 .517.06 1.02.174 1.5l-2.423.005a.75.75 0 0 1-.002-1.5zm8.751-2.506a.75.75 0 0 1 .743.648l.007.102v.75h-1.5v-.75a.75.75 0 0 1 .75-.75z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}