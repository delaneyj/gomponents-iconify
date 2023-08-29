package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceRecognitionLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.998 15v4h4v2h-6v-6h2Zm16 0v6h-6v-2h4v-4h2Zm-8-9v12h-2V6h2Zm-4 3v6h-2V9h2Zm8 0v6h-2V9h2Zm-8-6v2h-4v4h-2V3h6Zm12 0v6h-2V5h-4V3h6Z"/>`),
		g.Group(children),
	)
}