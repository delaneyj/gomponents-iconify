package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tunnel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M22.5 21V9.625M1.5 21V9.625M12 1.5v4m0-4c-2.939 0-5.366.703-7.144 1.474M12 1.5c2.939 0 5.366.703 7.144 1.474M0 22.5h24M5.5 21v-9c0-.46.048-.908.138-1.34M18.5 21v-9c0-.46-.048-.908-.138-1.34M12 5.5a6.473 6.473 0 0 0-4.136 1.485M12 5.5c1.571 0 3.012.557 4.136 1.485M22.5 16.5h-4m-17 0h4M4.856 2.974C2.7 3.91 1.5 4.942 1.5 4.942v4.683m3.356-6.65l3.008 4.01m0 0a6.498 6.498 0 0 0-2.226 3.675M1.5 9.625l4.138 1.035m12.724 0L22.5 9.625m-4.138 1.035a6.497 6.497 0 0 0-2.226-3.675m6.364 2.64V4.942s-1.2-1.033-3.356-1.968m-3.008 4.011l3.008-4.01"/>`),
		g.Group(children),
	)
}