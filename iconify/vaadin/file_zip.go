package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 15H5v-2.8l.7-2.2h2.4l.9 2.2V15zm4 0h-3v-3L9 9H7V8H5v1l-1 3v3H3V1h4v1h2v1H7v1h2v1h4v10zM10 4V1l3 3h-3z"/><path fill="currentColor" d="M5 6h2v1H5V6zm0-4h2v1H5V2zm0 2h2v1H5V4zm2 1h2v1H7V5zm0 2h2v1H7V7zm-1 5h2v2H6v-2z"/>`),
		g.Group(children),
	)
}