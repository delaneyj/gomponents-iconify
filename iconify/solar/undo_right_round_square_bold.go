package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoRightRoundSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.715 22 12c0-4.714 0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM15.76 6.873a.75.75 0 1 0-1.018 1.102l.84.776h-5.62a4.212 4.212 0 1 0 0 8.423H14.5a.75.75 0 0 0 0-1.5H9.962a2.712 2.712 0 1 1 0-5.423h5.62l-.84.776a.75.75 0 1 0 1.017 1.102l2.25-2.077a.75.75 0 0 0 0-1.102l-2.25-2.077Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}