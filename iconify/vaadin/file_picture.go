package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 0H2v16h12V4l-4-4zM9 5h4v10H3V1h6v4zm1-1V1l3 3h-3z"/><path fill="currentColor" d="M4 11.5V14h8v-1.7s.1-1.3-1.3-1.5c-1.3-.2-1.5.4-2.5.5c-.8 0-.6-1.3-2.2-1.3c-1.2 0-2 1.5-2 1.5zm8-3a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 12 8.5z"/>`),
		g.Group(children),
	)
}