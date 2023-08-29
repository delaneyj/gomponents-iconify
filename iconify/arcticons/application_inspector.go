package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationInspector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="21.5" r="17" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 38.5v5m8.727-24.003l1.507-1.484a1.758 1.758 0 0 0-2.491-2.48h-.001l-1.484 1.512a10.708 10.708 0 0 0-12.513 0L16.25 15.55a1.846 1.846 0 0 0-1.263-.532a1.732 1.732 0 0 0-1.234.51a1.755 1.755 0 0 0 0 2.48l1.512 1.483a10.758 10.758 0 0 0-2.021 6.285v2.209H34.76v-2.209a10.758 10.758 0 0 0-2.033-6.279Zm-11.557 4.06a1.699 1.699 0 1 1-1.699-1.699h0a1.699 1.699 0 0 1 1.7 1.699v0Zm7.361 1.698a1.699 1.699 0 1 1 1.699-1.698a1.699 1.699 0 0 1-1.699 1.698h0Z"/>`),
		g.Group(children),
	)
}