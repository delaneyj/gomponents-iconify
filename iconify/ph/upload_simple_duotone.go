package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M168 80H88l40-40Z" opacity=".2"/><path d="M224 152v56a16 16 0 0 1-16 16H48a16 16 0 0 1-16-16v-56a8 8 0 0 1 16 0v56h160v-56a8 8 0 0 1 16 0ZM80.61 83.06a8 8 0 0 1 1.73-8.72l40-40a8 8 0 0 1 11.32 0l40 40A8 8 0 0 1 168 88h-32v64a8 8 0 0 1-16 0V88H88a8 8 0 0 1-7.39-4.94ZM107.31 72h41.38L128 51.31Z"/></g>`),
		g.Group(children),
	)
}