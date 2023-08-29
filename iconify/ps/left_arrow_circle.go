package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 512q106 0 181-75t75-181t-75-181T256 0T75 75T0 256t75 181t181 75zm0-469q88 0 150.5 62.5T469 256t-62.5 150.5T256 469t-150.5-62.5T43 256t62.5-150.5T256 43zM149 262q3 0 0 0q0 3 2 5q0 2 3 2l42 64q12 17 30 6q18-10 7-30l-22-30h130q22 0 22-21t-22-21H211l22-30q11-20-7-30q-17-13-30 2l-42 64q0 2-3 2v11q-2 2-2 6z"/>`),
		g.Group(children),
	)
}