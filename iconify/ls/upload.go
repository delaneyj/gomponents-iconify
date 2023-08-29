package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M132 273h129v211c0 18 15 33 33 33h129c18 0 34-15 34-33V273h129c17 0 22-9 10-22L382 38c-6-7-15-11-24-11s-17 4-23 11L122 251c-12 13-8 22 10 22zM0 422v228c0 10 5 16 16 16h684c11 0 17-6 17-16V422c0-10-8-18-17-18h-65c-9 0-17 9-17 18v145H98V422c0-10-7-18-16-18H16c-9 0-16 9-16 18z"/>`),
		g.Group(children),
	)
}