package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotationOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRotationOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 4v36h36"/><path fill="#fff" d="M28 40c0-11.046-8.954-20-20-20v20h20Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRotationOne0)"/>`),
		g.Group(children),
	)
}