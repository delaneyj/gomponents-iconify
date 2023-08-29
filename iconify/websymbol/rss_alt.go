package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 89v824q0 36-26 62t-62 26H88q-36 0-62-26T0 913V89q0-36 26-62T88 1h824q36 0 62 26t26 62zM717 854h136q0-143-56-274T646.5 354.5T421 204t-274-56v135q154 0 285.5 76.5t208 208T717 854zm-240 0h136q0-126-62.5-233.5t-170-170T147 388v135q136 0 233 97q97 95 97 234zm-153-88q0-37-26-63t-63-26q-36 0-62 26t-26 63q0 36 26 62t62 26q37 0 63-26t26-62z"/>`),
		g.Group(children),
	)
}