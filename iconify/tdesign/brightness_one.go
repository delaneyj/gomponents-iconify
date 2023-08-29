package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .086l3.49 3.49h4.935V8.51l3.49 3.49l-3.49 3.49v4.934H15.49L12 23.914l-3.49-3.49H3.577V15.49L.086 12l3.49-3.49V3.575H8.51L12 .085Zm0 2.828L9.34 5.575H5.576V9.34L2.914 12l2.662 2.662v3.763h3.763L12 21.086l2.661-2.662h3.764v-3.763L21.085 12l-2.66-2.661V5.575H14.66L12 2.915ZM12.065 8A8.965 8.965 0 0 1 13 12a8.965 8.965 0 0 1-.936 4a4 4 0 0 0 0-8Zm-1.94-1.7a6 6 0 1 1 0 11.402l-1.314-.432l.823-1.113A6.964 6.964 0 0 0 11 12a6.964 6.964 0 0 0-1.367-4.157L8.81 6.731l1.314-.432Z"/>`),
		g.Group(children),
	)
}