package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDriveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaHardDriveFill0"><g id="evaHardDriveFill1"><path id="evaHardDriveFill2" fill="currentColor" d="m20.79 11.34l-3.34-6.68A3 3 0 0 0 14.76 3H9.24a3 3 0 0 0-2.69 1.66l-3.34 6.68a2 2 0 0 0-.21.9V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5.76a2 2 0 0 0-.21-.9ZM8 17a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm8 0h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2ZM5.62 11l2.72-5.45a1 1 0 0 1 .9-.55h5.52a1 1 0 0 1 .9.55L18.38 11Z"/></g></g>`),
		g.Group(children),
	)
}