package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Npo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="5.09" height="6.75" x="22.18" y="9.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.63" transform="rotate(-45 24.73 12.676)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.28 24.9l-3.62-3.62a1.63 1.63 0 0 0-2.31 0l-1.3 1.29a1.65 1.65 0 0 0 0 2.31l3.63 3.62"/><rect width="5.09" height="6.75" x="16.36" y="15.12" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.63" transform="rotate(-45 18.912 18.497)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.52 19.71l5.41 5.4"/><rect width="30" height="30" x="9" y="9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" transform="rotate(45 23.999 24)"/>`),
		g.Group(children),
	)
}