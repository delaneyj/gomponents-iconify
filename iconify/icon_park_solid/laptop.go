package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLaptop0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M8 8h32v20H8z"/><path d="M8 28L4 41h40l-4-13"/><path fill="#fff" d="M19.9 35h8.2l1.9 6H18l1.9-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLaptop0)"/>`),
		g.Group(children),
	)
}