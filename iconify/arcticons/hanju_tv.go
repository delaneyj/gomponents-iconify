package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HanjuTv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="39" height="30" x="4.5" y="9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><rect width="31.429" height="22" x="8.286" y="13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.095 18.397h7.425m-3.713 11.206V18.397m14.098 0l-3.712 11.206l-3.712-11.206M24.258 9l3.05-4.028M22.52 6.065L24.258 9"/>`),
		g.Group(children),
	)
}