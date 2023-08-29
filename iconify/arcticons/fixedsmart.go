package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fixedsmart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.767 14.212c-2.658-3.45-5.162-6.722-7.514-9.712H6.046c3.297 4.166 6.952 8.92 10.657 13.775Zm10.964 4.012A781.936 781.936 0 0 0 41.57 4.5h-6.874a458.283 458.283 0 0 1-6.952 9.737Zm.026 10.709l-4.09 4.089a2085.46 2085.46 0 0 1 7.489 10.452h7.053a1555.106 1555.106 0 0 0-10.452-14.541m-14.9.127A760.813 760.813 0 0 1 5.79 43.5h6.849a827.65 827.65 0 0 0 8.127-10.504Z"/>`),
		g.Group(children),
	)
}