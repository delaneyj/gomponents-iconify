package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M29 16.66a10.07 10.07 0 0 0 .25-2.24v-1a7.44 7.44 0 0 1-6.6-8.58A10.13 10.13 0 0 0 19 4.23a10.26 10.26 0 0 0-10.09 9.13A10 10 0 0 0 1 23.1c0 5.09 4.62 9.9 9.57 9.9h16.52c4.19 0 7.91-3.9 7.91-8.35a8.29 8.29 0 0 0-6-7.99Z" class="clr-i-solid--badged clr-i-solid-path-1--badged"/><circle cx="30" cy="6" r="5" fill="currentColor" class="clr-i-solid--badged clr-i-solid-path-2--badged clr-i-badge"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}