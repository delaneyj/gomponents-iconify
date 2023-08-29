package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CricketBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.07 4.93A10 10 0 0 0 4.93 19.07A10 10 0 0 0 19.07 4.93zM6.34 6.34a8 8 0 0 1 8.78-1.71l-.29.3l.71.71l.52-.53a9.53 9.53 0 0 1 .84.57L5.68 16.9a9.53 9.53 0 0 1-.57-.84l.53-.52l-.71-.71l-.29.29a8 8 0 0 1 1.7-8.78zm11.32 11.32a8 8 0 0 1-8.78 1.71l.29-.3l-.71-.71l-.52.53a9.53 9.53 0 0 1-.84-.57L18.32 7.1a9.53 9.53 0 0 1 .57.84l-.53.52l.71.71l.29-.29a8 8 0 0 1-1.7 8.78zm-6.37-2.12l.71.7l-1.41 1.42l-.71-.66zm2.83-2.83l.71.7l-1.42 1.42l-.7-.71zM17 9.88l.71.71L16.24 12l-.7-.71zm-4.29-1.42l-.71-.7l1.41-1.42l.71.71zm-2.83 2.83l-.71-.7l1.42-1.42l.7.71zm-2.83 2.83l-.71-.71L7.76 12l.7.71z"/>`),
		g.Group(children),
	)
}