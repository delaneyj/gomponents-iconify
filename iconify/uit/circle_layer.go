package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.979 2.068c-3.73-.524-7.18 2-7.849 5.672a5.985 5.985 0 0 0-3.086 4.658A5 5 0 0 0 7 22a5.007 5.007 0 0 0 4.607-3.06a5.963 5.963 0 0 0 4.635-3.076a6.985 6.985 0 0 0 5.69-5.893a6.996 6.996 0 0 0-5.953-7.903zM7 21a4 4 0 1 1 0-8a4 4 0 0 1 0 8zm4.907-3.089c.055-.296.093-.599.093-.911a5 5 0 0 0-5-5c-.312 0-.615.037-.912.092a4.997 4.997 0 1 1 5.82 5.82zm4.82-3.172l-.02.032a5.94 5.94 0 0 0 .288-1.529a5.997 5.997 0 0 0-5.746-6.237a5.973 5.973 0 0 0-1.985.258a5.996 5.996 0 1 1 7.464 7.476z"/>`),
		g.Group(children),
	)
}