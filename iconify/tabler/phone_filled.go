package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M9 3a1 1 0 0 1 .877.519l.051.11l2 5a1 1 0 0 1-.313 1.16l-.1.068l-1.674 1.004l.063.103a10 10 0 0 0 3.132 3.132l.102.062l1.005-1.672a1 1 0 0 1 1.113-.453l.115.039l5 2a1 1 0 0 1 .622.807L21 15v4c0 1.657-1.343 3-3.06 2.998C9.361 21.477 2.522 14.638 2 6a3 3 0 0 1 2.824-2.995L5 3h4z"/></g>`),
		g.Group(children),
	)
}