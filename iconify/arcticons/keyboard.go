package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.14 14.94v4.53h4.53v-4.53Zm6.8 0v4.53h4.53v-4.53Zm6.79 0v4.53h4.54v-4.53Zm6.8 0v4.53h4.53v-4.53Zm6.8 0v4.53h4.53v-4.53ZM8.14 21.73v4.54h4.53v-4.54Zm6.8 0v4.54h4.53v-4.54Zm6.79 0v4.54h4.54v-4.54Zm6.8 0v4.54h4.53v-4.54Zm6.8 0v4.54h4.53v-4.54Zm-27.19 6.8v4.53h4.53v-4.53Zm6.8 0v4.53h18.12v-4.53Zm20.39 0v4.53h4.53v-4.53Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 35.5v-23a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v23a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}