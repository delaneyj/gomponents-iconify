package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paysafecard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.417 17.24h17.166a5.584 5.584 0 0 1 5.585 5.585v15.09a5.584 5.584 0 0 1-5.585 5.585H15.417a5.584 5.584 0 0 1-5.585-5.584v-15.09a5.584 5.584 0 0 1 5.585-5.585Zm-2.905-1.91a11.508 11.508 0 0 1 19.382-7.696a11.636 11.636 0 0 1 3.594 7.695"/>`),
		g.Group(children),
	)
}