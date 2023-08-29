package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Peripage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.02 41.422H13.698A7.25 7.25 0 0 1 6.4 34.127V25.58c0-8.025 6.566-14.591 14.592-14.591h5.732c8.025 0 14.591 6.566 14.591 14.591v8.547a7.25 7.25 0 0 1-7.295 7.295Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.74 18.18c1.771-1.25 3.022-3.23 3.022-5.524c0-3.752-3.127-6.879-7.087-6.879c-3.335 0-6.15 2.293-6.88 5.316m-9.591 0c-.73-3.022-3.544-5.316-6.879-5.316c-3.96 0-7.087 3.127-7.087 6.88c0 2.292 1.25 4.273 3.022 5.523"/><circle cx="15.729" cy="26.622" r="5.941" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.789 34.856v-8.442"/><circle cx="31.988" cy="26.622" r="5.941" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.929 34.856v-8.442"/>`),
		g.Group(children),
	)
}