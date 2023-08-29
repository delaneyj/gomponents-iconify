package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitPullRequestSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M192 96a64 64 0 1 0-97 54.81v209.8a64 64 0 1 0 64 0V152a64.06 64.06 0 0 0 33-56Zm-64-32a32 32 0 1 1-32 32a32 32 0 0 1 32-32Zm-1 384a32 32 0 1 1 32-32a32 32 0 0 1-32 32Zm289-87.39V156a92.1 92.1 0 0 0-92-92h-35V9.93L201.14 96L289 182.07V128h35a28 28 0 0 1 28 28v204.61a64 64 0 1 0 64 0ZM384 448a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}