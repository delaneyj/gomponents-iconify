package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMuteTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 9.674V11h-1.326L12 9.326L10.326 11H9V9.674L10.674 8L9 6.326V5h1.326L12 6.674L13.674 5H15v1.326L13.326 8L15 9.674zM6.5 15a.504.504 0 0 1-.354-.146L2.292 11H.499a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5h1.793l3.854-3.854A.499.499 0 0 1 7 1.5v13a.5.5 0 0 1-.5.5z"/>`),
		g.Group(children),
	)
}