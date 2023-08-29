package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusStickerSaver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.018 2.503c-16.335.018-26.69 17.52-18.842 31.846L2.527 45.484l11.135-2.649c12.301 6.748 27.663.092 31.152-13.497C48.303 15.75 38.047 2.516 24.018 2.503Zm0 9.474v20.761"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.398 22.357l-10.38 10.381l-10.381-10.381M34.398 36.02H13.637"/>`),
		g.Group(children),
	)
}