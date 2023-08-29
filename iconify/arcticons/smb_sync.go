package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmbSync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="23.478" height="39" x="12.261" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.062"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.261 8.88h23.478M12.261 39.12h23.478M25.14 29.843a5.95 5.95 0 0 0 3.068-10.048m-5.34-1.639a5.962 5.962 0 0 0-3.076 10.055m7.768-7.768l-1.585 1.585v-4.356l4.348.008h0l-1.577 1.577l-1.186 1.186zm-7.112 7.112l1.577-1.577l.008 4.349h-4.357l1.585-1.585l1.187-1.187z"/>`),
		g.Group(children),
	)
}