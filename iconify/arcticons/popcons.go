package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Popcons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="16.5" cy="21.716" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="31.5" cy="21.716" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.58 27.636c-1.655 2.691-4.374 1.598-5.58.168c-1.206 1.43-3.925 2.523-5.58-.168"/><circle cx="24" cy="24" r="15.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}