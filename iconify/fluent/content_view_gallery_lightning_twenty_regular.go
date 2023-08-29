package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentViewGalleryLightningTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 7a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7Zm4 0H6v2h3V7Zm-3.5 4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4ZM5 13.5a.5.5 0 0 1 .5-.5h4c.015 0 .03 0 .044.002l-.404.86c-.022.046-.04.092-.057.138H5.5a.5.5 0 0 1-.5-.5ZM6 16a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h5v5.904l.96-2.042c.012-.027.026-.053.04-.08V4h2a2 2 0 0 1 2 2v1h1V6a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h5.971l.25-1H6Zm7.317-8h4.828a.5.5 0 0 1 .436.745L16.75 12h1.496a.75.75 0 0 1 .564 1.244l-4.824 5.508c-.504.576-1.442.085-1.257-.657L13.5 15h-3.003a.5.5 0 0 1-.452-.713l2.82-6A.5.5 0 0 1 13.317 8Z"/>`),
		g.Group(children),
	)
}