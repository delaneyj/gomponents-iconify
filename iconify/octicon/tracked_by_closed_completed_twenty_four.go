package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrackedByClosedCompletedTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 2.5A9.5 9.5 0 0 0 2.5 12a9.5 9.5 0 0 0 9.5 9.5a.75.75 0 0 1 0 1.5C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11a.75.75 0 0 1-1.5 0A9.5 9.5 0 0 0 12 2.5Z"/><path fill="currentColor" d="m13.759 17.48l3.728 3.314a.308.308 0 0 0 .513-.23V18h4.25a.75.75 0 0 0 0-1.5H18v-2.564a.308.308 0 0 0-.513-.23l-3.728 3.314a.307.307 0 0 0 0 .46Zm3.521-8.2a.749.749 0 1 0-1.06-1.06l-5.97 5.969l-2.47-2.469a.749.749 0 1 0-1.06 1.06l3 3a.749.749 0 0 0 1.06 0l6.5-6.5Z"/>`),
		g.Group(children),
	)
}