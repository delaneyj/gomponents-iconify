package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDrone0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m12 12l7 7m17 17l-7-7m7-17l-7 7M12 36l7-7"/><path fill="#fff" d="M19 19h10v10H19z"/><path d="M36 19a7 7 0 1 0-7-7m7 17a7 7 0 1 1-7 7m-17-7a7 7 0 1 0 7 7m-7-17a7 7 0 1 1 7-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDrone0)"/>`),
		g.Group(children),
	)
}