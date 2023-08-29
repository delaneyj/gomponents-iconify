package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowAltDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4v20.063l-4.281-4.282l-1.438 1.407L16 27.905l6.719-6.718l-1.438-1.407L17 24.063V4z"/>`),
		g.Group(children),
	)
}