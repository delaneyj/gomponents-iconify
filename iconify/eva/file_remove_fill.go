package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaFileRemoveFill0"><g id="evaFileRemoveFill1"><path id="evaFileRemoveFill2" fill="currentColor" d="m19.74 7.33l-4.44-5a1 1 0 0 0-.74-.33h-8A2.53 2.53 0 0 0 4 4.5v15A2.53 2.53 0 0 0 6.56 22h10.88A2.53 2.53 0 0 0 20 19.5V8a1 1 0 0 0-.26-.67ZM14 15h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Zm.71-7a.79.79 0 0 1-.71-.85V4l3.74 4Z"/></g></g>`),
		g.Group(children),
	)
}