package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Satellite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 17C2.97 12.794 2.97 6.118 7 2l15 15c-4.118 4.03-10.794 4.03-15 0Zm0 0c-3.295 0-6 2.95-6 6h12c0-1.139-.37-2.034-1-3m3-11l4-4l-4 4Zm5.5-8a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5Z"/>`),
		g.Group(children),
	)
}