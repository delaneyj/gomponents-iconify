package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchingPathsDownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM12 5.75a.75.75 0 0 1 .75.75v6A3.25 3.25 0 0 0 16 15.75h.19l-.22-.22a.75.75 0 1 1 1.06-1.06l1.5 1.5a.75.75 0 0 1 0 1.06l-1.5 1.5a.75.75 0 1 1-1.06-1.06l.22-.22H16a4.746 4.746 0 0 1-4-2.187a4.746 4.746 0 0 1-4 2.187h-.19l.22.22a.75.75 0 1 1-1.06 1.06l-1.5-1.5a.75.75 0 0 1 0-1.06l1.5-1.5a.75.75 0 0 1 1.06 1.06l-.22.22H8a3.25 3.25 0 0 0 3.25-3.25v-6a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}