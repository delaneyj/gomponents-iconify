package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageWithCurl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#CCD6DD" d="M28 32a4 4 0 0 1-4 4H4c-2.209 0-4-1.875-4-8V4a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v28z"/><path fill="#E1E8ED" d="M31 36H4c4 0 4-8 4-8a4 4 0 0 1 4-4h20c2.209 0 4 2 4 4c0 0 .25 8-5 8z"/><path fill="#99AAB5" d="M24 7a1 1 0 0 1-1 1H5a1 1 0 0 1 0-2h18a1 1 0 0 1 1 1zm0 4a1 1 0 0 1-1 1H5a1 1 0 0 1 0-2h18a1 1 0 0 1 1 1zm0 4a1 1 0 0 1-1 1H5a1 1 0 0 1 0-2h18a1 1 0 0 1 1 1zm0 4a1 1 0 0 1-1 1H5a1 1 0 1 1 0-2h18a1 1 0 0 1 1 1z"/>`),
		g.Group(children),
	)
}