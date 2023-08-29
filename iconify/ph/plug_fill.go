package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M237.66 77.66L203.31 112l26.35 26.34a8 8 0 0 1-11.32 11.32l-6.34-6.35l-53 53a40 40 0 0 1-56.57 0l-15.68-15.74l-49.09 49.09a8 8 0 0 1-11.32-11.32l49.09-49.09l-15.71-15.71a40 40 0 0 1 0-56.57l53-53l-6.35-6.34a8 8 0 0 1 11.32-11.32L144 52.69l34.34-34.35a8 8 0 1 1 11.32 11.32L155.31 64L192 100.69l34.34-34.35a8 8 0 0 1 11.32 11.32Z"/>`),
		g.Group(children),
	)
}