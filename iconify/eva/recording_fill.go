package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaRecordingFill0"><g id="evaRecordingFill1"><path id="evaRecordingFill2" fill="currentColor" d="M18 8a4 4 0 0 0-4 4a3.91 3.91 0 0 0 .56 2H9.44a3.91 3.91 0 0 0 .56-2a4 4 0 1 0-4 4h12a4 4 0 0 0 0-8Z"/></g></g>`),
		g.Group(children),
	)
}