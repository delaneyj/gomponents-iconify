package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11a1 1 0 0 1-1-.999V7a5.002 5.002 0 0 1 8.532-3.542a5.078 5.078 0 0 1 1.307 2.293a1 1 0 1 1-1.937.501v-.003a3.057 3.057 0 0 0-.786-1.379A3.002 3.002 0 0 0 9 7v3a1 1 0 0 1-.999 1H8zm5.5 3.5a1.5 1.5 0 1 0-3 0c0 .443.195.836.5 1.11v1.392A1 1 0 0 0 12 18h.001A1 1 0 0 0 13 17v-1.39c.305-.274.5-.667.5-1.11z" opacity=".5"/><path fill="currentColor" d="M17 9H7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3zm-4 6.61V17a1 1 0 0 1-.999 1H12a1 1 0 0 1-1-.999V15.61a1.49 1.49 0 0 1-.5-1.11a1.5 1.5 0 1 1 3 0c0 .443-.195.836-.5 1.11z"/>`),
		g.Group(children),
	)
}