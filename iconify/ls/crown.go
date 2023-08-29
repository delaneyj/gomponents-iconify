package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M384 185c-44 0-80-36-80-80c0-45 36-80 80-80s80 35 80 80c0 44-36 80-80 80zM65 152c35 0 63 29 63 64c0 36-28 64-63 64c-36 0-65-28-65-64c0-35 29-64 65-64zm640 0c35 0 63 29 63 64c0 36-28 64-63 64c-36 0-65-28-65-64c0-35 29-64 65-64zM135 594L82 304c16-3 31-12 43-22c33 30 75 59 115 59c48 0 87-74 114-136c9 3 20 6 30 6s21-3 30-6c27 62 66 136 114 136c40 0 82-29 115-59c12 10 27 19 43 22l-53 290H135zm-11 25h520v49H124v-49z"/>`),
		g.Group(children),
	)
}