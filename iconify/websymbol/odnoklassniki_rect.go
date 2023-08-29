package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OdnoklassnikiRect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 89v824q0 36-26 62t-62 26H88q-36 0-62-26T0 913V89q0-36 26-62T88 1h824q36 0 62 26t26 62zM705 305q0-87-59-150T500 92t-146 63t-59 150t59 150t146 63t146-63t59-150zm22 279q0-26-17.5-44.5T667 521q-15 0-28 8q-70 36-140 36q-69 0-138-36q-16-7-28-7q-25 0-42.5 18.5T273 585q0 68 147 97L303 803q-18 18-18 44t17.5 44.5T345 910t43-18l112-116l112 116q18 18 43 18t42.5-18.5T715 847t-18-44L580 682q147-30 147-98zM585 305q0 36-24.5 62T500 393t-60.5-26t-24.5-62t24.5-62t60.5-26t60.5 26t24.5 62z"/>`),
		g.Group(children),
	)
}