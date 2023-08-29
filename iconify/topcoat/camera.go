package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M14.41 22.58c0 3.359 2.73 6.09 6.09 6.09c3.359 0 6.09-2.73 6.09-6.09s-2.73-6.09-6.09-6.09a6.095 6.095 0 0 0-6.09 6.09zM3.5 36.5h34c2.63 0 3-.37 3-3v-23c0-2.462-.38-3-3-3h-10c0-2.57-.42-3-3-3h-8c-2.55 0-3 .48-3 3h-10c-2.58 0-3 .692-3 3v23c0 2.6.38 3 3 3zm7.64-13.92c0-5.17 4.19-9.359 9.36-9.359s9.359 4.189 9.359 9.359s-4.189 9.359-9.359 9.359s-9.36-4.189-9.36-9.359z"/>`),
		g.Group(children),
	)
}