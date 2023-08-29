package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceptionHotelBell(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 5.5c5.344 0 8.705 2.8 9.376 7.63c.163 1.176 1.124 2.37 2.124 3.12V17H.5v-.75c1-.75 1.96-1.944 2.124-3.12C3.295 8.3 6.656 5.5 12 5.5Zm0 0v-3m-12 17h24m-14-17h4"/>`),
		g.Group(children),
	)
}