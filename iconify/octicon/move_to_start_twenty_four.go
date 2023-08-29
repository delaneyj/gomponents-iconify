package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveToStartTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.78 18.78a.749.749 0 0 1-1.06 0l-6.25-6.25a.749.749 0 0 1 0-1.06l6.25-6.25a.749.749 0 1 1 1.06 1.06l-4.969 4.97H22.25a.75.75 0 0 1 0 1.5H7.811l4.969 4.97a.749.749 0 0 1 0 1.06ZM2.75 3.75a.75.75 0 0 1 .75.75v15a.75.75 0 0 1-1.5 0v-15a.75.75 0 0 1 .75-.75Z"/>`),
		g.Group(children),
	)
}