package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600c222.411 0 416.39-121.104 520.02-300.879L908.79 777.612C847.217 884.405 732.127 956.47 600 956.47c-196.873 0-356.47-159.597-356.47-356.47S403.127 243.53 600 243.53c84.387 0 161.732 29.521 222.729 78.589L665.186 434.18L1200 612.524V53.613l-174.17 123.926C917.124 67.952 766.553 0 600 0z"/>`),
		g.Group(children),
	)
}