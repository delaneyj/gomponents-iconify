package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalAltThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 16h-4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5zM6 22H3v-5h3v5zM22.5 2h-4a.5.5 0 0 0-.5.5v20a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-20a.5.5 0 0 0-.5-.5zM22 22h-3V3h3v19zm-7.5-12h-4a.5.5 0 0 0-.5.5v12a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-12a.5.5 0 0 0-.5-.5zM14 22h-3V11h3v11z"/>`),
		g.Group(children),
	)
}