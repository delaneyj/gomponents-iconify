package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kwai(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.85 42.5h13.478a5.391 5.391 0 0 0 5.39-5.391v-7.764a5.391 5.391 0 0 0-5.39-5.391H20.85a5.391 5.391 0 0 0-5.39 5.39v7.765a5.391 5.391 0 0 0 5.39 5.391Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="18.328" cy="13.261" r="7.761"/><circle cx="32.904" cy="14.206" r="6.815"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.46 30.594l-4.483-2.588c-1.198-.692-2.696.173-2.696 1.556v7.353c0 1.383 1.498 2.248 2.696 1.556l4.482-2.588"/>`),
		g.Group(children),
	)
}