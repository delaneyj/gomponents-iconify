package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KatalonIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#19D89F" d="M256 261.407H127.998v130.725H256z"/><path d="M128.001 0L0 130.725v130.726h128.001l127.958-130.726V0z"/>`),
		g.Group(children),
	)
}