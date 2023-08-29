package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrentLocationSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.75 2a.75.75 0 0 0-1.5 0v1.784a8.252 8.252 0 0 0-7.466 7.466H2a.75.75 0 0 0 0 1.5h1.784a8.252 8.252 0 0 0 7.466 7.466V22a.75.75 0 0 0 1.5 0v-1.784a8.252 8.252 0 0 0 7.466-7.466H22a.75.75 0 0 0 0-1.5h-1.784a8.252 8.252 0 0 0-7.466-7.466V2Zm-4.5 10a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}