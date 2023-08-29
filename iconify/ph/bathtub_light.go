package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BathtubLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 98h-26v-2a6 6 0 0 0-6-6h-64a6 6 0 0 0-6 6v2H62V52a14 14 0 0 1 14-14a14.47 14.47 0 0 1 14.12 11.19a6 6 0 1 0 11.76-2.38A26.32 26.32 0 0 0 76 26a26 26 0 0 0-26 26v46H24a14 14 0 0 0-14 14v32a54.06 54.06 0 0 0 54 54h2v18a6 6 0 0 0 12 0v-18h100v18a6 6 0 0 0 12 0v-18h2a54.06 54.06 0 0 0 54-54v-32a14 14 0 0 0-14-14Zm-90 4h52v36h-52Zm92 42a42 42 0 0 1-42 42H64a42 42 0 0 1-42-42v-32a2 2 0 0 1 2-2h106v34a6 6 0 0 0 6 6h64a6 6 0 0 0 6-6v-34h26a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}