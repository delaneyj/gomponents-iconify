package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContrastOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3.055a9.001 9.001 0 0 0 0 17.89V3.055Zm2 0V4.84l2.184-1.26A8.946 8.946 0 0 0 13 3.054Zm4.255 1.638L13 7.149v4.042l6.693-3.864a9.047 9.047 0 0 0-2.438-2.634Zm3.284 4.455L13 13.5v4.042l7.95-4.59a9.093 9.093 0 0 0-.411-3.804Zm-.326 6.539L13 19.85v1.094a9.008 9.008 0 0 0 7.213-5.258ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Z"/>`),
		g.Group(children),
	)
}