package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11.536 10.536a5 5 0 0 0-7.071-7.071m7.07 7.07a5 5 0 0 1-7.071-7.071m7.072 7.072L4.464 3.464M12 17c0 1.28.488 2.56 1.464 3.535A4.984 4.984 0 0 0 17 22m-5-5a5 5 0 1 1 10 0m-10 0h10m0 0c0 1.28-.488 2.56-1.465 3.535"/>`),
		g.Group(children),
	)
}