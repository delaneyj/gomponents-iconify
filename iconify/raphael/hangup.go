package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hangup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.563 10.494c-7.35-7.35-19.265-7.348-26.612 0c-1.795 1.797-.246 6.84-.246 6.84a.917.917 0 0 0 1.067.615l6.9-1.605c.45-.105.82-.57.82-1.033v-3.685c0-.463.38-.842.842-.842h8.285c.464 0 .843.38.843.842v3.47c0 .463.374.908.83.987l7.634 1.316a.787.787 0 0 0 .926-.692s.51-4.418-1.287-6.214zm-11.3 3.578h-3.5v4.39h-2.625l4.363 7.556l4.364-7.556h-2.6v-4.39z"/>`),
		g.Group(children),
	)
}