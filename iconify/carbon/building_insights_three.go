package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingInsightsThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H16a2.002 2.002 0 0 0-2 2v10H4a2.002 2.002 0 0 0-2 2v14h28V4a2.002 2.002 0 0 0-2-2ZM9 28v-7h4v7Zm19 0H15v-8a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v8H4V16h12V4h12Z"/><path fill="currentColor" d="M18 8h2v2h-2zm6 0h2v2h-2zm-6 6h2v2h-2zm6 0h2v2h-2zm-6 6h2v2h-2zm6 0h2v2h-2zM2 10h5v2H2zm8-8h2v5h-2zM4 5.414L5.414 4L9 7.585L7.585 9z"/>`),
		g.Group(children),
	)
}