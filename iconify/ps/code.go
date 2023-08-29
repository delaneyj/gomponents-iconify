package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M312 3q-8-4-16 .5T286 18L169 404q-4 18 15 25q8 4 16-.5t10-14.5L325 28q7-18-13-25zm45 320q14 14 30 0l106-107l-106-107q-16-14-30 0q-16 16 0 30l77 77l-77 77q-12 14 0 30zm-250 0q16 14 30 0q16-16 0-30l-75-77l77-77q16-14 0-30q-14-14-30 0L3 216z"/>`),
		g.Group(children),
	)
}