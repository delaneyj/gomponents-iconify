package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.5 5.5h-5.281L9.51 4.26A2.5 2.5 0 0 0 7.34 3H4.5A2.5 2.5 0 0 0 2 5.5V15a2.5 2.5 0 0 0 2.5 2.5h11A2.5 2.5 0 0 0 18 15V8a2.5 2.5 0 0 0-2.5-2.5Zm-11 11A1.5 1.5 0 0 1 3 15V5.5A1.5 1.5 0 0 1 4.5 4h2.84a1.5 1.5 0 0 1 1.302.756l.852 1.492a.5.5 0 0 0 .435.252H15.5A1.5 1.5 0 0 1 17 8v7a1.5 1.5 0 0 1-1.5 1.5h-11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}