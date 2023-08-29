package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogOutFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLogOutFill0"><g id="evaLogOutFill1"><path id="evaLogOutFill2" fill="currentColor" d="M7 6a1 1 0 0 0 0-2H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2H6V6Zm13.82 5.42l-2.82-4a1 1 0 0 0-1.39-.24a1 1 0 0 0-.24 1.4L18.09 11H10a1 1 0 0 0 0 2h8l-1.8 2.4a1 1 0 0 0 .2 1.4a1 1 0 0 0 .6.2a1 1 0 0 0 .8-.4l3-4a1 1 0 0 0 .02-1.18Z"/></g></g>`),
		g.Group(children),
	)
}