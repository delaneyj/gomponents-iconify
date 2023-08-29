package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserNurse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.94 22H3.06a1 1 0 0 1-.994-1.108a9.995 9.995 0 0 1 19.868 0A1 1 0 0 1 20.94 22Z" opacity=".5"/><path fill="currentColor" d="m12.708 18.307l4.706-4.715a10.001 10.001 0 0 0-10.833.003l4.712 4.712A1 1 0 0 0 12 18.6a1.002 1.002 0 0 0 .708-.293Z" opacity=".25"/><path fill="currentColor" d="M11.995 14a6 6 0 1 1 6-6a6.007 6.007 0 0 1-6 6Z" opacity=".25"/><path fill="currentColor" d="M6.09 9a5.993 5.993 0 0 0 11.82 0Z"/>`),
		g.Group(children),
	)
}