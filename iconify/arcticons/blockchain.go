package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blockchain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.169 41.33L6.67 26.831a4.016 4.016 0 0 1 0-5.662L21.169 6.67a4.016 4.016 0 0 1 5.662 0L41.33 21.169a4.016 4.016 0 0 1 0 5.662L26.831 41.33a4.004 4.004 0 0 1-5.662 0ZM24 42.502V24m-13.86-6.301L24 24m13.86-6.301L24 24"/>`),
		g.Group(children),
	)
}