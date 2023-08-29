package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceRecognitionFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.998 3v18h-18V3h18Zm-8 3h-2v12h2V6Zm-4 3h-2v6h2V9Zm8 0h-2v6h2V9Z"/>`),
		g.Group(children),
	)
}