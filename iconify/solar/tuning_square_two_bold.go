package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningSquareTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 17.25a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5ZM15.25 8a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Z"/><path fill-rule="evenodd" d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Zm8 6.75a2.75 2.75 0 1 1 0-5.5a2.75 2.75 0 0 1 0 5.5ZM16.75 8a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0Zm-3.5 8a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 0 1.5h-5a.75.75 0 0 1-.75-.75ZM10 7.25a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1 0-1.5h5ZM4.25 16a.75.75 0 0 1 .75-.75h1a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1-.75-.75ZM19 7.25a.75.75 0 0 1 0 1.5h-1a.75.75 0 0 1 0-1.5h1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}