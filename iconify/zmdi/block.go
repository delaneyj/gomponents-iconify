package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Block(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zM43 216q0 59 36 105L318 81q-46-36-105-36q-70 0-120 50T43 216zm170 171q71 0 121-50t50-121q0-59-36-105L109 351q46 36 104 36z"/>`),
		g.Group(children),
	)
}