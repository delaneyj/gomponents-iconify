package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLoopCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd"><path d="M8 7a2.5 2.5 0 0 1 2.5-2.5h5.1a1 1 0 0 1 .702.288l4.4 4.333a1 1 0 0 1 .298.712V17a2.5 2.5 0 0 1-2.5 2.5H12a1 1 0 1 1 0-2h6.5a.5.5 0 0 0 .5-.5v-6.167h-2.4a2 2 0 0 1-2-2V6.5h-4.1a.5.5 0 0 0-.5.5v1.5a1 1 0 0 1-2 0V7Zm8.6.888l.96.945h-.96v-.945Z"/><path d="M11.049 11.678a2.193 2.193 0 1 0-2.058 3.873a2.193 2.193 0 0 0 2.058-3.873Zm-4.732-.031a4.193 4.193 0 1 1 2.674 6.033l-1.676 3.155a1 1 0 0 1-1.767-.938l1.677-3.155a4.195 4.195 0 0 1-.908-5.095Z"/></g><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}