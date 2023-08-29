package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerGapLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M134 32v32a6 6 0 0 1-12 0V32a6 6 0 0 1 12 0Zm90 90h-32a6 6 0 0 0 0 12h32a6 6 0 0 0 0-12Zm-46.5 47a6 6 0 0 0-8.5 8.5l22.63 22.62a6 6 0 0 0 8.48-8.48ZM128 186a6 6 0 0 0-6 6v32a6 6 0 0 0 12 0v-32a6 6 0 0 0-6-6Zm-49.5-17l-22.62 22.64a6 6 0 1 0 8.48 8.48L87 177.5a6 6 0 1 0-8.5-8.5ZM70 128a6 6 0 0 0-6-6H32a6 6 0 0 0 0 12h32a6 6 0 0 0 6-6Zm-5.64-72.12a6 6 0 0 0-8.48 8.48L78.5 87a6 6 0 1 0 8.5-8.5Z"/>`),
		g.Group(children),
	)
}