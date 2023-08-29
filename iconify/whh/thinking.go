package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thinking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 304q0 42-18.5 78T955 443q5 20 5 37q0 66-47 113t-113 47q-35 0-69-16q-26 37-67 58.5T576 704q-53 0-97.5-27T409 604q-38 36-89 36q-53 0-90.5-37.5T192 512q-80 0-136-56T0 320t56-136t136-56q41 0 79 17q24-64 81-104.5T480 0q50 0 94 21.5T650 80q14-36 46-58t72-22q53 0 90.5 37.5T896 128q0 1-.5 3.5t-.5 3.5q57 16 93 62.5t36 106.5zM192 896q27 0 45.5 18.5t18.5 45t-19 45.5t-45.5 19t-45-19t-18.5-45t18.5-45t45.5-19zm160-192q40 0 68 28t28 68t-28 68t-68 28t-68-28t-28-68t28-68t68-28z"/>`),
		g.Group(children),
	)
}