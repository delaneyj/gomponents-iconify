package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderNew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.5 2.5V4h5.697l-1-1.5H1.5ZM0 2v10.75C0 13.44.56 14 1.25 14h6a.75.75 0 0 0 0-1.5H1.5v-7h13v1.75a.75.75 0 0 0 1.5 0v-2C16 4.56 15.44 4 14.75 4H8.998a1.054 1.054 0 0 0-.034-.055l-1.667-2.5A1 1 0 0 0 6.465 1H1a1 1 0 0 0-1 1Zm13 13a.75.75 0 0 1-.75-.75v-1.5h-1.5a.75.75 0 0 1 0-1.5h1.5v-1.5a.75.75 0 0 1 1.5 0v1.5h1.5a.75.75 0 0 1 0 1.5h-1.5v1.5A.75.75 0 0 1 13 15Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}