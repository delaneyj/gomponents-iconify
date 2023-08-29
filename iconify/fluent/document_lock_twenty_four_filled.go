package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentLockTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2v6a2 2 0 0 0 2 2h6v10a2 2 0 0 1-2 2h-6.05c.033-.162.05-.329.05-.5v-5a2.5 2.5 0 0 0-2-2.45V14a3.5 3.5 0 0 0-6-2.45V4a2 2 0 0 1 2-2h6Zm1.5.5V8a.5.5 0 0 0 .5.5h5.5l-6-6ZM4 15h-.5A1.5 1.5 0 0 0 2 16.5v5A1.5 1.5 0 0 0 3.5 23h6a1.5 1.5 0 0 0 1.5-1.5v-5A1.5 1.5 0 0 0 9.5 15H9v-1a2.5 2.5 0 0 0-5 0v1Zm1.5-1a1 1 0 1 1 2 0v1h-2v-1Zm2 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/>`),
		g.Group(children),
	)
}