package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextboxRotateNinetyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.25 21A2.75 2.75 0 0 0 21 18.25V5.75A2.75 2.75 0 0 0 18.25 3H5.75A2.75 2.75 0 0 0 3 5.75v12.5A2.75 2.75 0 0 0 5.75 21h12.5Zm-5.75-6.75a.75.75 0 0 1-1.493.102L11 14.25v-7.5a.75.75 0 0 1 1.493-.102l.007.102v7.5Zm-4-7.5v10.5a.75.75 0 0 1-1.493.102L7 17.25V6.75a.75.75 0 0 1 1.493-.102l.007.102Zm8 10.5a.75.75 0 0 1-1.493.102L15 17.25V6.75a.75.75 0 0 1 1.493-.102l.007.102v10.5Z"/>`),
		g.Group(children),
	)
}