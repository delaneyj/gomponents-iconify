package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargeBlueDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2.565 18.121a3 3 0 0 1 0-4.242L13.879 2.565a3 3 0 0 1 4.242 0l11.314 11.314a3 3 0 0 1 0 4.242L18.12 29.435a3 3 0 0 1-4.242 0L2.565 18.121Zm1.414-2.828a1 1 0 0 0 0 1.414l11.314 11.314a1 1 0 0 0 1.414 0l11.314-11.314a1 1 0 0 0 0-1.414L16.707 3.979a1 1 0 0 0-1.414 0L3.979 15.293Z"/>`),
		g.Group(children),
	)
}