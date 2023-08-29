package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConfoundedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSConfoundedFace0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" stroke-linecap="round" d="m32 17l-3 3l3 3m-16-6l3 3l-3 3m-1 8l3 3l3-3l3 3l3-3l3 3l3-3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSConfoundedFace0)"/>`),
		g.Group(children),
	)
}