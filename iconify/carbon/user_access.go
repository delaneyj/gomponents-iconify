package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserAccess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 30h-2v-5a3.003 3.003 0 0 0-3-3H7a3.003 3.003 0 0 0-3 3v5H2v-5a5.006 5.006 0 0 1 5-5h4a5.006 5.006 0 0 1 5 5zM9 10a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0-2a5 5 0 1 0 5 5a5 5 0 0 0-5-5zm21 4a1.992 1.992 0 0 0-.512.074l-1.055-1.055a3.879 3.879 0 0 0 0-4.038l1.055-1.055a2.034 2.034 0 1 0-1.414-1.414L27.02 5.567a3.879 3.879 0 0 0-4.038 0l-1.055-1.055a2.034 2.034 0 1 0-1.414 1.414l1.055 1.055a3.879 3.879 0 0 0 0 4.038l-1.055 1.055a2.034 2.034 0 1 0 1.414 1.414l1.055-1.055a3.879 3.879 0 0 0 4.038 0l1.055 1.055A1.996 1.996 0 1 0 30 12zm-7-3a2 2 0 1 1 2 2a2.002 2.002 0 0 1-2-2z"/>`),
		g.Group(children),
	)
}