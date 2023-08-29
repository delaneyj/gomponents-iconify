package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crrowd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.434 37.44a3.008 3.008 0 1 0 0 6.016h23.133a3.008 3.008 0 1 0 0-6.016Zm17.742 6.06v-5.762m-.296-22.39h8.216a14.532 14.532 0 1 0 0 7.371H29.88"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.88 15.348h.012a6.971 6.971 0 0 0-5.896-3.258h0a6.943 6.943 0 1 0 5.899 10.629h-.015"/>`),
		g.Group(children),
	)
}