package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shopping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M644 480v46H203L126 73H0V28h165l77 452h402zm73-453l-69 400H287L217 27h500zM334 150V73h-62l14 77h48zm46 0h63V73h-63v77zm109-77v77h64V73h-64zm110 77h50l14-77h-64v77zm-265 46h-41l11 64h30v-64zm109 64v-64h-63v64h63zm46 0h64v-64h-64v64zm152-64h-42v64h31zM334 306h-22l13 75h9v-75zm46 75h63v-75h-63v75zm173-75h-64v75h64v-75zm46 0v75h10l14-75h-24zM254 553c32 0 56 25 56 56c0 32-24 57-56 57c-31 0-56-25-56-57c0-31 25-56 56-56zm343 0c32 0 56 25 56 56c0 32-24 57-56 57s-56-25-56-57c0-31 24-56 56-56z"/>`),
		g.Group(children),
	)
}