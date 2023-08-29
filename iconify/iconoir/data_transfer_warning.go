package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTransferWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 4v1m0 4v1m10 2V4m0 0l3 3m-3-3l-3 3m6 9v2m0 4.01l.01-.011M7 14v6m0 0l3-3m-3 3l-3-3"/>`),
		g.Group(children),
	)
}