package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="3" fill="currentColor"/><path fill="currentColor" d="M12 13c-2.761 0-5 1.79-5 4h10c0-2.21-2.239-4-5-4z"/><rect width="18" height="18" x="3" y="3" rx="3"/></g>`),
		g.Group(children),
	)
}