package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V3H6zm2 2h16v22H8V5zm8 4.72v6.468A2.902 2.902 0 0 0 15 16c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3v-6.72l2.75.69l.5-1.94l-4-1L16 9.72zM15 18c.564 0 1 .436 1 1c0 .564-.436 1-1 1c-.564 0-1-.436-1-1c0-.564.436-1 1-1z"/>`),
		g.Group(children),
	)
}