package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M84 28v82H2v164h128V28H84zM48 233v-82h36v82H48zm92-123h46v164h-46V110zm0-82h46v46h-46V28zm189 82H206v164h77v36h-77v46h123V110zm-77 123v-82h31v82h-31zm210-123H340v164h81v36h-81v46h122V110zm-41 123h-35v-82h35v82z"/>`),
		g.Group(children),
	)
}