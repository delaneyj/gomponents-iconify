package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmileyStickerLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 30 199.52a6 6 0 0 0 2.48-1.49L224 160.46a6 6 0 0 0 1.49-2.49A102.06 102.06 0 0 0 128 26Zm86.47 127l-61.42 61.43A89.91 89.91 0 1 1 214.47 153ZM82 108a10 10 0 1 1 10 10a10 10 0 0 1-10-10Zm92 0a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm-.81 47c-9.92 17.16-26.39 27-45.19 27s-35.27-9.84-45.19-27a6 6 0 0 1 10.38-6c7.84 13.54 20.2 21 34.81 21s27-7.46 34.81-21a6 6 0 1 1 10.38 6Z"/>`),
		g.Group(children),
	)
}