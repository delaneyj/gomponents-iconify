package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowArrowFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m245.66 85.66l-32 32a8 8 0 0 1-11.32-11.32L220.69 88H192c-1.66 0-41.35.77-48.11 41.31C139 158.6 121.75 172 108.13 178a71.12 71.12 0 0 1-27.49 6H79a32 32 0 1 1 0-16h1.33c1.62 0 40.77.81 47.79-41.32C132.82 98.44 150 85 163.65 78.73A74.49 74.49 0 0 1 192 72h28.69l-18.35-18.34a8 8 0 0 1 11.32-11.32l32 32a8 8 0 0 1 0 11.32Z"/>`),
		g.Group(children),
	)
}