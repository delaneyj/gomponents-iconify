package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Platte(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPlatte0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M24 44c5.96 0 2.336-8.864 6-13c3.126-3.53 14-1.914 14-7c0-11.046-8.954-20-20-20S4 12.954 4 24s8.954 20 20 20Z"/><path d="M28 17a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm-12 4a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm1 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPlatte0)"/>`),
		g.Group(children),
	)
}