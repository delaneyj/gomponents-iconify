package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Harrypotter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1000 886H27q-43-40-15-66L472 15q17-15 41-15t40 15l460 805q29 27-13 66zM82 818h239q-50-40-79.5-97T207 598zm193-239.5q0 89.5 58.5 156T479 814V344q-87 12-145.5 78.5t-58.5 156zM375 305q49-25 104-31V123zm173-180v149q54 6 102 31zm0 219v470q87-13 145.5-79.5t58.5-156t-58.5-156T548 344zm271 256q-4 65-34 121.5T706 818h238z"/>`),
		g.Group(children),
	)
}