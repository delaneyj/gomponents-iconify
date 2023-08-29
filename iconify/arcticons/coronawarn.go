package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coronawarn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.8a7.5 7.5 0 0 1 0-15a7.32 7.32 0 0 1 5.3 2.2l9.9-9.9A21.9 21.9 0 0 0 24 2.8a21.5 21.5 0 1 0 15.2 36.7l-9.9-9.9a7.32 7.32 0 0 1-5.3 2.2Z"/><circle cx="24" cy="11.3" r="1.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 16.8v-3.6"/><circle cx="14.8" cy="15.1" r="1.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.7 19l-2.6-2.5"/><circle cx="11" cy="24.3" r="1.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.5 24.3h-3.6"/><circle cx="14.8" cy="33.6" r="1.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.7 29.6l-2.6 2.6"/><circle cx="24" cy="37.4" r="1.9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 31.8v3.7"/>`),
		g.Group(children),
	)
}