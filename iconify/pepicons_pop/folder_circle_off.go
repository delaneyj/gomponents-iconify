package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M18.5 8h-4.991l-.565-.988A3 3 0 0 0 10.34 5.5H7.5a3 3 0 0 0-3 3V18a3 3 0 0 0 3 3h11a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3Zm-11 11a1 1 0 0 1-1-1V8.5a1 1 0 0 1 1-1h2.84a1 1 0 0 1 .868.504l.852 1.492a1 1 0 0 0 .869.504H18.5a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-11Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}