package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookRect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 89v824q0 36-26 62t-62 26H88q-36 0-62-26T0 913V89q0-36 26-62T88 1h824q36 0 62 26t26 62zM813 249l20-118q-43-15-143-15q-79 0-123 58q-26 34-26 119v78h-79v115h79v400h149V486h118l9-115H690v-90q0-42 59-42q31 0 64 10z"/>`),
		g.Group(children),
	)
}