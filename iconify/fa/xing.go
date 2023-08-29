package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M597 667q-10 18-257 456q-27 46-65 46H36q-21 0-31-17t0-36l253-448q1 0 0-1L97 388q-12-22-1-37q9-15 32-15h239q40 0 66 45zm806-642q11 16 0 37L875 996v1l336 615q11 20 1 37q-10 15-32 15H941q-42 0-66-45L536 997q18-32 531-942q25-45 64-45h241q22 0 31 15z"/>`),
		g.Group(children),
	)
}