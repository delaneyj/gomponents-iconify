package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m362.7 347l-29.9-44.85L302.9 347l-42.6-28.4l37-55.5l-41.3-27.532l-41.3 27.532l37 55.5l-42.6 28.4l-29.9-44.85L149.3 347l-42.6-28.4l54.04-81.06l69.66-46.44V128h51.2v63.1l69.66 46.44l54.04 81.06l-42.6 28.4zm98.1-91C460.8 98.93 289.594.227 153.369 78.763s-136.225 275.94 0 354.474S460.8 413.071 460.8 256zM127.711 477.547c-170.281-98.17-170.281-344.925 0-443.094S512 59.662 512 256S297.993 575.716 127.711 477.547z"/>`),
		g.Group(children),
	)
}