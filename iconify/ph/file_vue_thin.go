package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileVueThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m83.77 153.35l-20 56a4 4 0 0 1-7.54 0l-20-56a4 4 0 0 1 7.54-2.7L60 196.11l16.23-45.46a4 4 0 0 1 7.54 2.7ZM208 156a4 4 0 0 0 0-8h-32a4 4 0 0 0-4 4v56a4 4 0 0 0 4 4h32a4 4 0 0 0 0-8h-28v-20h20a4 4 0 0 0 0-8h-20v-20Zm-64-8a4 4 0 0 0-4 4v38a14 14 0 0 1-28 0v-38a4 4 0 0 0-8 0v38a22 22 0 0 0 44 0v-38a4 4 0 0 0-4-4Zm68-60v24a4 4 0 0 1-8 0V92h-52a4 4 0 0 1-4-4V36H56a4 4 0 0 0-4 4v72a4 4 0 0 1-8 0V40a12 12 0 0 1 12-12h96a4 4 0 0 1 2.83 1.17l56 56A4 4 0 0 1 212 88Zm-13.66-4L156 41.65V84Z"/>`),
		g.Group(children),
	)
}