package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaRecordingOutline0"><g id="evaRecordingOutline1"><path id="evaRecordingOutline2" fill="currentColor" d="M18 8a4 4 0 0 0-4 4a3.91 3.91 0 0 0 .56 2H9.44a3.91 3.91 0 0 0 .56-2a4 4 0 1 0-4 4h12a4 4 0 0 0 0-8ZM4 12a2 2 0 1 1 2 2a2 2 0 0 1-2-2Zm14 2a2 2 0 1 1 2-2a2 2 0 0 1-2 2Z"/></g></g>`),
		g.Group(children),
	)
}