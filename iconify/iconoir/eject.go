package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 14l-.592-.46A.75.75 0 0 0 5 14.75V14Zm14 0v.75a.75.75 0 0 0 .592-1.21L19 14Zm-14 .75h14v-1.5H5v1.5Zm5.619-9.196L4.408 13.54l1.184.92l6.21-7.985l-1.183-.92Zm8.973 7.986l-6.21-7.986l-1.185.921l6.211 7.986l1.184-.921Zm-7.79-7.065a.25.25 0 0 1 .395 0l1.184-.92a1.749 1.749 0 0 0-2.762 0l1.184.92ZM5 17.25a.75.75 0 0 0 0 1.5v-1.5Zm14 1.5a.75.75 0 0 0 0-1.5v1.5Zm-14 0h14v-1.5H5v1.5Z"/>`),
		g.Group(children),
	)
}