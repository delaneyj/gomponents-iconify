package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535l17.072-17.07C19.07 2 16.713 2 12 2C7.286 2 4.929 2 3.464 3.464Z" clip-rule="evenodd"/><path d="M3.465 20.536C4.929 22 7.286 22 12 22s7.071 0 8.536-1.464C22 19.07 22 16.714 22 12c0-4.714 0-7.07-1.464-8.535L3.465 20.535Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}