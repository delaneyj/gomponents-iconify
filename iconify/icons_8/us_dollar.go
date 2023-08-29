package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsDollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 4v2.03c-2.77.203-5 2.526-5 5.345c0 2.584 1.864 4.788 4.406 5.25l.594.125V24h-.625A3.368 3.368 0 0 1 11 20.625V19H9v1.625C9 23.575 11.426 26 14.375 26H15v2h2v-2h.625C20.573 26 23 23.574 23 20.625c0-2.584-1.863-4.788-4.406-5.25L17 15.062v-7.03a3.36 3.36 0 0 1 3 3.343V13h2v-1.625c0-2.82-2.23-5.142-5-5.344V4h-2zm0 4.03v6.658l-.25-.032a3.314 3.314 0 0 1-2.75-3.28a3.361 3.361 0 0 1 3-3.345zm2 9.095l1.25.22a3.313 3.313 0 0 1 2.75 3.28A3.369 3.369 0 0 1 17.625 24H17v-6.875z"/>`),
		g.Group(children),
	)
}