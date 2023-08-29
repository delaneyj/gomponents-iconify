package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.75 2c0 2.9 2.35 5.25 5.25 5.25a.75.75 0 0 1 0 1.5a6.737 6.737 0 0 1-5.25-2.507V18a4.75 4.75 0 1 1-1.5-3.464V2h1.5Zm-1.5 16a3.25 3.25 0 1 0-6.5 0a3.25 3.25 0 0 0 6.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}