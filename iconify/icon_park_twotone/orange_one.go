package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTOrangeOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M34.15 34.85c8.202-8.202 8.202-21.498 0-29.7L4.453 34.85c8.201 8.2 21.498 8.2 29.699 0Z"/><path d="m19.301 20l14.85 14.85M19.302 20v21.213m0-21.213l20.506.707m-.053-5.71c1.664 6.88-.189 14.437-5.56 19.808c-5.372 5.372-12.93 7.225-19.809 5.56m-.034-15.415l8.485-8.486"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTOrangeOne0)"/>`),
		g.Group(children),
	)
}