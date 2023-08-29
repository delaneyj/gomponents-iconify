package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VinylOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.746 1.976a.75.75 0 0 1-.672.82a9.251 9.251 0 1 0 10.13 10.13a.75.75 0 1 1 1.493.149c-.54 5.433-5.122 9.675-10.697 9.675c-5.937 0-10.75-4.813-10.75-10.75c0-5.574 4.243-10.157 9.676-10.696a.75.75 0 0 1 .82.672Zm2.81-.123a.75.75 0 0 1 .669-.11a10.774 10.774 0 0 1 7.033 7.032a.75.75 0 1 1-1.431.45a9.278 9.278 0 0 0-5.077-5.683V12a3.75 3.75 0 1 1-1.5-3V2.457a.75.75 0 0 1 .306-.605Zm-.306 10.148a2.25 2.25 0 1 0-4.5 0a2.25 2.25 0 0 0 4.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}