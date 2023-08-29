package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTApplicationOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M41 14L24 4L7 14v20l17 10l17-10V14Z"/><path stroke-linecap="round" d="M16 18.998L23.993 24l8.002-5.002M24 24v9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTApplicationOne0)"/>`),
		g.Group(children),
	)
}