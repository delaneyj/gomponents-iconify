package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DsNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.77 42.5v-37M42.5 19.81v8.38H32.32c-2.32 0-4.19-1.88-4.19-4.19c0-1.16.47-2.2 1.23-2.96c.75-.76 1.8-1.23 2.96-1.23H42.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 6.93v34.14c0 .79-.64 1.43-1.43 1.43H6.93c-.79 0-1.43-.64-1.43-1.43V6.93c0-.79.64-1.43 1.43-1.43h34.14c.79 0 1.43.64 1.43 1.43Z"/>`),
		g.Group(children),
	)
}