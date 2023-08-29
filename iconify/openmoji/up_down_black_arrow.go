package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpDownBlackArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3f3f3f" d="m52.509 48.279l-4.1-3.7l-8.6 9.2V17.478l8.6 9.2l4.1-3.7l-15.8-16.7l-15.8 16.7l4.1 3.7l8.6-9.2v36.301l-8.6-9.2l-4.1 3.7l15.8 16.7l15.8-16.7z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m52.509 48.279l-4.1-3.7l-8.6 9.2V17.478l8.6 9.2l4.1-3.7l-15.8-16.7l-15.8 16.7l4.1 3.7l8.6-9.2v36.301l-8.6-9.2l-4.1 3.7l15.8 16.7l15.8-16.7z"/>`),
		g.Group(children),
	)
}