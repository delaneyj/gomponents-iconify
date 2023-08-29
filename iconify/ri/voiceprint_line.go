package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VoiceprintLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7h2v10H5V7Zm-4 3h2v4H1v-4Zm8-8h2v18H9V2Zm4 2h2v18h-2V4Zm4 3h2v10h-2V7Zm4 3h2v4h-2v-4Z"/>`),
		g.Group(children),
	)
}