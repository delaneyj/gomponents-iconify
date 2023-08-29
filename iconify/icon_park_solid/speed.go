package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSpeed0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M34.023 6.69A19.908 19.908 0 0 0 24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20c0-3.627-.966-7.03-2.654-9.962"/><path fill="#fff" d="M31.95 16.05S28.562 25.095 27 26.657A4 4 0 0 1 21.343 21c1.562-1.562 10.607-4.95 10.607-4.95Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSpeed0)"/>`),
		g.Group(children),
	)
}