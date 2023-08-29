package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendTwiceSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM17 14.75a.75.75 0 0 1 0 1.5h-4.928l1.408 1.174a.75.75 0 0 1-.96 1.152l-3-2.5a.75.75 0 0 1 0-1.152l3-2.5a.75.75 0 0 1 .96 1.152l-1.409 1.174H17Zm.75-6.25a.75.75 0 0 0-.75-.75h-4.928l1.408-1.174a.75.75 0 0 0-.96-1.152l-3 2.5a.75.75 0 0 0 0 1.152l3 2.5a.75.75 0 0 0 .96-1.152L12.071 9.25H17a.75.75 0 0 0 .75-.75ZM7.75 6a.75.75 0 0 0-1.5 0v12a.75.75 0 0 0 1.5 0V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}