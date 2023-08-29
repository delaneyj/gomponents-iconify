package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureInPictureExitSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 7a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h5Zm-5 4.5V8h5a2 2 0 0 0 2-2V4h3.5A2.5 2.5 0 0 1 15 6.5v5a2.5 2.5 0 0 1-2.5 2.5h-8A2.5 2.5 0 0 1 2 11.5Zm8.5-.5a.5.5 0 0 0 0 1h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-1 0v.793l-1.646-1.647a.5.5 0 0 0-.708.708L11.293 11H10.5Z"/>`),
		g.Group(children),
	)
}