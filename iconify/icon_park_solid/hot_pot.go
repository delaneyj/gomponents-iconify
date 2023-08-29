package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotPot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHotPot0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M24 11V4m8 7V7m-16 4V7"/><path fill="#fff" stroke="#fff" d="M44 17H4c0 5.586 3.578 10.503 9 13.365C16.156 32.03 19.936 33 24 33s7.844-.97 11-2.635c5.422-2.862 9-7.78 9-13.365Z"/><path stroke="#fff" d="M10.467 39h27.066M13 30.365L9 44m26-13.635L39 44"/><path stroke="#000" d="M20 25h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHotPot0)"/>`),
		g.Group(children),
	)
}