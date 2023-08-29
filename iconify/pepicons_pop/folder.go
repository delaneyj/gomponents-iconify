package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.5 5h-4.991l-.565-.988A3 3 0 0 0 7.34 2.5H4.5a3 3 0 0 0-3 3V15a3 3 0 0 0 3 3h11a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3Zm-11 11a1 1 0 0 1-1-1V5.5a1 1 0 0 1 1-1h2.84a1 1 0 0 1 .868.504l.852 1.492A1 1 0 0 0 9.93 7h5.57a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-11Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}