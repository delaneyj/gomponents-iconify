package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRingOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M24 11.273A3.636 3.636 0 1 0 24 4a3.636 3.636 0 0 0 0 7.273ZM24 44a3.636 3.636 0 1 0 0-7.273A3.636 3.636 0 0 0 24 44ZM7.636 27.636a3.636 3.636 0 1 0 0-7.272a3.636 3.636 0 0 0 0 7.272Zm32.728 0a3.636 3.636 0 1 0 0-7.273a3.636 3.636 0 0 0 0 7.273Z"/><path stroke-linecap="round" d="M32.734 10.16a16.45 16.45 0 0 1 5.107 5.106m-.001 17.468a16.45 16.45 0 0 1-5.106 5.106m-17.467 0a16.45 16.45 0 0 1-5.107-5.106m0-17.468a16.45 16.45 0 0 1 5.107-5.106"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRingOne0)"/>`),
		g.Group(children),
	)
}