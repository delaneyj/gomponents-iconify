package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zalo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.466 19.266h6.716l-6.716 10.137h6.716"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.517 4.502l-29.011.035c-8.624 7.4-11.224 17.418-3.53 31.911c.487 1.708-.359 3.42-2 5.128a11.17 11.17 0 0 0 6.864-1.265c13.983 6.27 21.919 2.18 29.644-2.733l.035-31.074a2 2 0 0 0-1.997-2.002h-.005h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.63 26.791c0 1.442-1.22 2.612-2.724 2.612s-2.724-1.17-2.724-2.612v-1.698c0-1.442 1.22-2.612 2.724-2.612s2.724 1.17 2.724 2.612m9.13-2.61h0a2.714 2.714 0 0 1 2.714 2.714v1.494a2.714 2.714 0 0 1-2.713 2.714h0a2.714 2.714 0 0 1-2.714-2.714v-1.494a2.714 2.714 0 0 1 2.714-2.714Zm-9.13 6.92v-6.921m2.681-3.527v9.143c0 .722.585 1.307 1.306 1.307h.392"/>`),
		g.Group(children),
	)
}