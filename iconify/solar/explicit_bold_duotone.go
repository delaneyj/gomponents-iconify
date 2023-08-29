package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExplicitBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 22c-4.714 0-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12s0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22Z" opacity=".5"/><path d="M8.25 8c0-.966.784-1.75 1.75-1.75h5a.75.75 0 0 1 0 1.5h-5a.25.25 0 0 0-.25.25v3.25H15a.75.75 0 0 1 0 1.5H9.75V16c0 .138.112.25.25.25h5a.75.75 0 0 1 0 1.5h-5A1.75 1.75 0 0 1 8.25 16V8Z"/></g>`),
		g.Group(children),
	)
}