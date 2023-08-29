package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtVoice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.535 24.91v-2.65c0-.468.285-.888.72-1.06l3.651-1.391V8.806a1.32 1.32 0 0 1 2.62 0v.68c5.492 3.681 14.284 7.152 22.286 7.152h5.832a2 2 0 0 1 2 2v.89h0a1.91 1.91 0 0 1 1.89 1.921V25.7a1.9 1.9 0 0 1-1.95 1.84v.921a2 2 0 0 1-2 2h-2.2l-2.311 7.743a.89.89 0 0 1-.87.65H30.15c-.09.01-.18.01-.27 0a1 1 0 0 1-.65-1.18l1.29-4.322h-.62a.838.838 0 0 1-.27 0a1 1 0 0 1-.64-1.18l.39-1.21a45.352 45.352 0 0 0-17.824 6.811v.62c.01.084.01.168 0 .25h0a1.311 1.311 0 1 1-2.611-.25V27.392L5.295 26a1.14 1.14 0 0 1-.76-1.09Zm4.371 2.451v-7.542"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.381 30.962c2.654-.332 5.308-.553 7.982-.5"/>`),
		g.Group(children),
	)
}