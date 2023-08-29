package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalXLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212.24 195.76a6 6 0 1 1-8.48 8.48L184 184.48l-19.76 19.76a6 6 0 0 1-8.48-8.48L175.52 176l-19.76-19.76a6 6 0 0 1 8.48-8.48L184 167.52l19.76-19.76a6 6 0 0 1 8.48 8.48L192.48 176ZM160 126a6 6 0 0 0 6-6V72a6 6 0 0 0-12 0v48a6 6 0 0 0 6 6Zm40 0a6 6 0 0 0 6-6V32a6 6 0 0 0-12 0v88a6 6 0 0 0 6 6Zm-80-20a6 6 0 0 0-6 6v88a6 6 0 0 0 12 0v-88a6 6 0 0 0-6-6Zm-40 40a6 6 0 0 0-6 6v48a6 6 0 0 0 12 0v-48a6 6 0 0 0-6-6Zm-40 40a6 6 0 0 0-6 6v8a6 6 0 0 0 12 0v-8a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}