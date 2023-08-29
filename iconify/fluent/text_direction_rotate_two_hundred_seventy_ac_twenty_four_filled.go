package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRotateTwoHundredSeventyAcTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M19.707 6.707a1 1 0 0 0 0-1.414l-2-2a1 1 0 0 0-1.414 0l-2 2a1 1 0 0 0 1.414 1.414L16 6.414V20a1 1 0 1 0 2 0V6.414l.293.293a1 1 0 0 0 1.414 0zm-9-1.414a1 1 0 0 1-1.414 1.414L9 6.414V8a1 1 0 0 1-2 0V6.414l-.293.293a1 1 0 0 1-1.414-1.414l2-2a1 1 0 0 1 1.414 0l2 2zM4 12a1 1 0 1 0 0 2h9.25a1 1 0 1 0 0-2H9v-1a1 1 0 1 0-2 0v1H4zm0 8a1 1 0 1 0 2 0v-3h.25c.895 0 1.87.184 2.586.642C9.49 18.061 10 18.745 10 20a1 1 0 1 0 2 0c0-1.945-.864-3.26-2.086-4.042C8.756 15.216 7.354 15 6.25 15H5a1 1 0 0 0-1 1v4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}