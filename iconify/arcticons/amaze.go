package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amaze(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.746 19.305l1.502 3.879l3.879 1.502l-3.879 1.501l-1.502 3.879l-1.502-3.879l-3.879-1.501l3.879-1.502l1.502-3.879zm7.647 7.154l1.006 2.6L34 30.066l-2.601 1.007l-1.006 2.601l-1.007-2.601l-2.601-1.007l2.601-1.007l1.007-2.6zm-13.508 1.465l.806 2.081l2.08.805l-2.08.805l-.806 2.08l-.805-2.08L14 30.81l2.08-.805l.805-2.081z"/>`),
		g.Group(children),
	)
}