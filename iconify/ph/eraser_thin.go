package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EraserThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m222.15 83.22l-41.38-41.37a20 20 0 0 0-28.28 0L33.85 160.48a20 20 0 0 0 0 28.29l30.06 30.06a4 4 0 0 0 2.83 1.17H216a4 4 0 0 0 0-8h-94.34l100.49-100.49a20 20 0 0 0 0-28.29ZM110.34 212H68.4l-28.89-28.89a12 12 0 0 1 0-17L96 109.65L154.35 168Zm106.15-106.15L160 162.34L101.65 104l56.49-56.48a12 12 0 0 1 17 0l41.37 41.37a12 12 0 0 1-.02 16.96Z"/>`),
		g.Group(children),
	)
}