package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimalisticMagniferZoomInOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 2.75a8.75 8.75 0 1 0 0 17.5a8.75 8.75 0 0 0 0-17.5ZM1.25 11.5c0-5.66 4.59-10.25 10.25-10.25S21.75 5.84 21.75 11.5S17.16 21.75 11.5 21.75S1.25 17.16 1.25 11.5ZM11.5 8.25a.75.75 0 0 1 .75.75v1.75H14a.75.75 0 0 1 0 1.5h-1.75V14a.75.75 0 0 1-1.5 0v-1.75H9a.75.75 0 0 1 0-1.5h1.75V9a.75.75 0 0 1 .75-.75Zm7.97 11.22a.75.75 0 0 1 1.06 0l2 2a.75.75 0 1 1-1.06 1.06l-2-2a.75.75 0 0 1 0-1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}