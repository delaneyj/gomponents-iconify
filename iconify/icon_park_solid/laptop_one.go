package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLaptopOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M8 11a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v21H8V11Z"/><path fill="#fff" d="M4 32h40v2a6 6 0 0 1-6 6H10a6 6 0 0 1-6-6v-2Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLaptopOne0)"/>`),
		g.Group(children),
	)
}