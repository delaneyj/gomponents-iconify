package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardholderBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535Z" opacity=".5"/><path d="M8 15.25a.75.75 0 0 0 0 1.5h8a.75.75 0 0 0 0-1.5H8Z"/><path fill-rule="evenodd" d="M17 10.25H5a.75.75 0 0 0 0 1.5h14a.75.75 0 0 0 0-1.5h-2Z" clip-rule="evenodd"/><path d="M13 6h-2c-1.886 0-2.828 0-3.414.586C7 7.172 7 8.114 7 10v.25h10V10c0-1.886 0-2.828-.586-3.414C15.828 6 14.886 6 13 6Z" opacity=".7"/></g>`),
		g.Group(children),
	)
}