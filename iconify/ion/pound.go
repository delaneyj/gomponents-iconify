package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M125.2 352.3H32v-54.6h101.2l13.1-83.3H47v-54.6h107.2L176 32h63.9l-21.8 127.7h105.6L345.5 32h63.1l-21.8 127.7H480v54.6H378.1l-12.3 83.3H465v54.6H358.5L336 480h-63.1l21.8-127.7H188.3L166.5 480h-63.1l21.8-127.7zm84.2-138L197 297.7h105.6l12.3-83.3H209.4z" fill="currentColor"/>`),
		g.Group(children),
	)
}