package circle_flags

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ItTwentyThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><mask id="circleFlagsIt230" width="512" height="512" x="0" y="0" maskUnits="userSpaceOnUse"><circle cx="256" cy="256" r="256" fill="#fff"/></mask></defs><g mask="url(#circleFlagsIt230)"><path fill="#333" d="M0 0h256l64 256l-64 256H0V0Z"/><path fill="#d80027" d="M256 0h256v512H256V0Z"/></g>`),
		g.Group(children),
	)
}