package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayStore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.027 24.052L8.765 39.885h0C9.285 41.65 10.945 43 12.918 43c.83 0 1.557-.208 2.18-.623h0l18.17-10.279l-8.24-8.046Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.95 20.262h0l-7.786-4.464l-8.137 8.254l8.24 8.046l7.788-4.36c1.35-.727 2.284-2.18 2.284-3.738c-.104-1.558-1.038-3.011-2.388-3.738ZM8.765 8.115c-.104.311-.104.727-.104 1.142v29.59c0 .415 0 .727.104 1.142l16.262-15.937L8.765 8.115Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.027 24.052l8.137-8.254L15.202 5.623C14.58 5.208 13.75 5 12.918 5C10.945 5 9.18 6.35 8.765 8.115h0l16.262 15.937Z"/>`),
		g.Group(children),
	)
}