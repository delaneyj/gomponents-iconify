package mi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.806 12.919a4 4 0 1 0-3.787-6.586a1 1 0 0 0 1.49 1.334a1.992 1.992 0 0 1 1.695-.657A2 2 0 0 1 18 11H3a1 1 0 1 0 0 2h15c.276 0 .546-.028.806-.081zM5 10h5.516a2.5 2.5 0 1 0-1.88-4.167a1 1 0 0 0 1.491 1.334A.5.5 0 1 1 10.5 8H5a1 1 0 0 0 0 2zm11.5 4H8a1 1 0 1 0 0 2h8.5a.5.5 0 1 1-.373.833a1 1 0 1 0-1.49 1.334A2.5 2.5 0 1 0 16.517 14H16.5z"/>`),
		g.Group(children),
	)
}