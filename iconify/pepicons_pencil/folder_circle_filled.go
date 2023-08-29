package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilFolderCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M18.5 8.5h-5.281l-.709-1.24A2.5 2.5 0 0 0 10.34 6H7.5A2.5 2.5 0 0 0 5 8.5V18a2.5 2.5 0 0 0 2.5 2.5h11A2.5 2.5 0 0 0 21 18v-7a2.5 2.5 0 0 0-2.5-2.5Zm-11 11A1.5 1.5 0 0 1 6 18V8.5A1.5 1.5 0 0 1 7.5 7h2.84a1.5 1.5 0 0 1 1.302.756l.852 1.492a.5.5 0 0 0 .435.252H18.5A1.5 1.5 0 0 1 20 11v7a1.5 1.5 0 0 1-1.5 1.5h-11Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilFolderCircleFilled0)"/></g>`),
		g.Group(children),
	)
}