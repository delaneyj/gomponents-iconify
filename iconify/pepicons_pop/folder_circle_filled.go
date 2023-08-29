package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopFolderCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M18.5 8h-4.991l-.565-.988A3 3 0 0 0 10.34 5.5H7.5a3 3 0 0 0-3 3V18a3 3 0 0 0 3 3h11a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3Zm-11 11a1 1 0 0 1-1-1V8.5a1 1 0 0 1 1-1h2.84a1 1 0 0 1 .868.504l.852 1.492a1 1 0 0 0 .869.504H18.5a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-11Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopFolderCircleFilled0)"/></g>`),
		g.Group(children),
	)
}