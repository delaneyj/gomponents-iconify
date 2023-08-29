package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaHeadphonesFill0"><g id="evaHeadphonesFill1"><path id="evaHeadphonesFill2" fill="currentColor" d="M12 2A10.2 10.2 0 0 0 2 12.37V17a4 4 0 1 0 4-4a3.91 3.91 0 0 0-2 .56v-1.19A8.2 8.2 0 0 1 12 4a8.2 8.2 0 0 1 8 8.37v1.19a3.91 3.91 0 0 0-2-.56a4 4 0 1 0 4 4v-4.63A10.2 10.2 0 0 0 12 2Z"/></g></g>`),
		g.Group(children),
	)
}