package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrackPrev(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M10.31 19.802l-6.56-6.249c-1-.799-1-2.307 0-3.106l6.564-6.252c.67-.48 1.686 0 1.686.805v4l5.394-4.808C18.063 3.714 19 4.192 19 5v14c0 .812-.936 1.285-1.602.809L12 15v4c0 .816-1.02 1.281-1.69.802z" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}