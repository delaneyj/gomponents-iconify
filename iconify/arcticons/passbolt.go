package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passbolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31.541" height="31.541" x="8.23" y="8.23" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.352" transform="rotate(45 24 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 28.359V24"/><circle cx="7.859" cy="24" r="4.359" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.491 24H12.218"/>`),
		g.Group(children),
	)
}