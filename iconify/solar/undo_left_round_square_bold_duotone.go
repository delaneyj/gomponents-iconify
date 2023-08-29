package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoLeftRoundSquareBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2c4.714 0 7.071 0 8.535 1.464C22 4.93 22 7.286 22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z" opacity=".5"/><path fill-rule="evenodd" d="M9.301 6.915a.75.75 0 0 1-.042 1.06l-.84.775h5.62a4.212 4.212 0 0 1 0 8.423H9.5a.75.75 0 1 1 0-1.5h4.539a2.712 2.712 0 0 0 0-5.423h-5.62l.84.776a.75.75 0 1 1-1.018 1.102l-2.25-2.077a.75.75 0 0 1 0-1.102l2.25-2.077a.75.75 0 0 1 1.06.043Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}