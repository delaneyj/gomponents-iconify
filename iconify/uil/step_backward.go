package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.28 3.43a3.23 3.23 0 0 0-3.29 0L8 8.84V6a3 3 0 0 0-6 0v12a3 3 0 0 0 6 0v-2.84l9 5.37a3.28 3.28 0 0 0 1.68.47a3.24 3.24 0 0 0 1.61-.43a3.38 3.38 0 0 0 1.72-3V6.42a3.38 3.38 0 0 0-1.73-2.99ZM6 18a1 1 0 0 1-2 0V6a1 1 0 0 1 2 0Zm14-.42a1.4 1.4 0 0 1-.71 1.25a1.23 1.23 0 0 1-1.28 0l-9.33-5.6a1.45 1.45 0 0 1 0-2.46L18 5.19a1.23 1.23 0 0 1 .67-.19a1.29 1.29 0 0 1 .62.17A1.4 1.4 0 0 1 20 6.42Z"/>`),
		g.Group(children),
	)
}