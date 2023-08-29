package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropletOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaDropletOffFill0"><g id="evaDropletOffFill1"><path id="evaDropletOffFill2" fill="currentColor" d="M19 16.14A7.73 7.73 0 0 0 17.34 8l-4.56-4.69a1 1 0 0 0-.71-.31a1 1 0 0 0-.72.3L8.74 5.92ZM6 8.82a7.73 7.73 0 0 0 .64 9.9A7.44 7.44 0 0 0 11.92 21a7.34 7.34 0 0 0 4.64-1.6Zm14.71 10.47l-16-16a1 1 0 0 0-1.42 1.42l16 16a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`),
		g.Group(children),
	)
}