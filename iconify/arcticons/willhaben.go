package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Willhaben(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.504 13.131l-2.184 8.738l-2.185-8.738l-2.184 8.738l-2.185-8.738m11.469 0v8.738m2.73-8.738v8.738h4.369m2.185-8.738v8.738h4.369m-14.02 2.641l-3.742 6.474m9.492-6.474l-5.127 8.919m-4.365-2.444c-2 3.464 2.33 5.964 4.33 2.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".979" d="M42.5 24.51h-27a7.01 7.01 0 0 1-7.01-7.01h0a7.01 7.01 0 0 1 7.01-7.01h27"/>`),
		g.Group(children),
	)
}