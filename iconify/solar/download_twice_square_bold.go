package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadTwiceSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM9.25 7a.75.75 0 0 0-1.5 0v4.928L6.576 10.52a.75.75 0 0 0-1.152.96l2.5 3a.75.75 0 0 0 1.152 0l2.5-3a.75.75 0 0 0-1.152-.96L9.25 11.929V7Zm6.25-.75a.75.75 0 0 1 .75.75v4.928l1.174-1.408a.75.75 0 0 1 1.152.96l-2.5 3a.75.75 0 0 1-1.152 0l-2.5-3a.75.75 0 0 1 1.152-.96l1.174 1.409V7a.75.75 0 0 1 .75-.75Zm-9.5 10a.75.75 0 0 0 0 1.5h12a.75.75 0 0 0 0-1.5H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}