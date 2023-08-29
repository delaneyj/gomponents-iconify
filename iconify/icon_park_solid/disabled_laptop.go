package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisabledLaptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDisabledLaptop0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M23 8.999H11a3 3 0 0 0-3 3v21h32v-7"/><path fill="#fff" stroke-linejoin="round" d="M4 32.999h40v2a6 6 0 0 1-6 6H10a6 6 0 0 1-6-6v-2Z"/><circle cx="36" cy="13" r="6"/><path stroke-linecap="round" stroke-linejoin="round" d="m32 9l8 8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDisabledLaptop0)"/>`),
		g.Group(children),
	)
}