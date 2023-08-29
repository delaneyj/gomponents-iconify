package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosedCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilLockClosedCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M10.5 15.5a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15 10.5h-5A3.5 3.5 0 0 0 6.5 14v3a3.5 3.5 0 0 0 3.5 3.5h5a3.5 3.5 0 0 0 3.5-3.5v-3a3.5 3.5 0 0 0-3.5-3.5ZM7.5 14a2.5 2.5 0 0 1 2.5-2.5h5a2.5 2.5 0 0 1 2.5 2.5v3a2.5 2.5 0 0 1-2.5 2.5h-5A2.5 2.5 0 0 1 7.5 17v-3Z" clip-rule="evenodd"/><path d="M10 11a.5.5 0 0 1-1 0V7.5a3.5 3.5 0 1 1 7 0V11a.5.5 0 0 1-1 0V7.5a2.5 2.5 0 0 0-5 0V11Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilLockClosedCircleFilled0)"/></g>`),
		g.Group(children),
	)
}