package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiMap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="24" cy="28.154" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.03" ry="3.105"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 20.609a7.597 7.597 0 0 0-7.093 5.766a10.111 10.111 0 0 0-.106 3.034c.488 4.06 4.223 7.83 7.199 11.001c2.977-3.172 6.711-6.942 7.2-11a10.111 10.111 0 0 0-.107-3.035A7.597 7.597 0 0 0 24 20.609ZM4.5 17.85a23.663 23.663 0 0 1 39-.004"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.236 22.929a16.118 16.118 0 0 1 27.528 0"/>`),
		g.Group(children),
	)
}