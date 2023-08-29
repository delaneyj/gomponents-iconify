package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnaTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3v1c-.007 2.46-.91 4.554-2.705 6.281M12 12c-3.328 1.99-5 4.662-5.008 8.014v1m10.008 0v-1c0-1.44-.315-2.755-.932-3.944M12 12c-1.903-1.138-3.263-2.485-4.082-4.068M8 4h9M7 20h10M12 8h4m-8 8h8M3 3l18 18"/>`),
		g.Group(children),
	)
}