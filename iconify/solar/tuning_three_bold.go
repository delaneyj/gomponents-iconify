package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningThreeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 8.75a.75.75 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v3a.75.75 0 0 1-.75.75ZM4 12a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm6 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm8 2a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm1.25-4a.75.75 0 0 0 1.5 0V5a.75.75 0 0 0-1.5 0v5ZM4 13.25a.75.75 0 0 0-.75.75v5a.75.75 0 0 0 1.5 0v-5a.75.75 0 0 0-.75-.75ZM11.25 19a.75.75 0 0 0 1.5 0v-3a.75.75 0 0 0-1.5 0v3Zm8.75.75a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 1 1.5 0v1a.75.75 0 0 1-.75.75ZM3.25 5a.75.75 0 0 1 1.5 0v1a.75.75 0 0 1-1.5 0V5Z"/>`),
		g.Group(children),
	)
}