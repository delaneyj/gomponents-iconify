package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Location(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M43 0v13.166C27.944 16.03 16.03 27.944 13.166 43H0v14h13.166C16.03 72.056 27.944 83.97 43 86.834V100h14V86.834C72.056 83.97 83.97 72.056 86.834 57H100V43H86.834C83.97 27.944 72.056 16.03 57 13.166V0H43zm7 22.5A27.425 27.425 0 0 1 77.5 50A27.425 27.425 0 0 1 50 77.5A27.425 27.425 0 0 1 22.5 50A27.425 27.425 0 0 1 50 22.5z" color="currentColor"/>`),
		g.Group(children),
	)
}