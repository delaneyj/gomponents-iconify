package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortPantsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16.8h6.966a.6.6 0 0 0 .596-.53l1.36-11.6a.6.6 0 0 0-.596-.67H3.659a.6.6 0 0 0-.597.656l1.387 14.8a.6.6 0 0 0 .597.544H11.4a.6.6 0 0 0 .6-.6V12"/>`),
		g.Group(children),
	)
}