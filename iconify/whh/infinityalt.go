package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Infinityalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 632q-119 0-203.5-84.5T448 344q0-66-47-113t-113-47t-113 47t-47 113t47 113t113 47q67 0 114-48q22 63 64 114q-78 62-178 62q-119 0-203.5-84.5T0 344t84.5-203.5T288 56t203.5 84.5T576 344q0 66 47 113t113 47t113-47t47-113t-47-113t-113-47q-67 0-114 48q-22-63-64-114q78-62 178-62q119 0 203.5 84.5T1024 344t-84.5 203.5T736 632z"/>`),
		g.Group(children),
	)
}