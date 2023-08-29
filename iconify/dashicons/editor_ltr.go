package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 2H5.5C3 2 1 4 1 6.5S3 11 5.5 11H6v6c0 .5.5 1 1 1s1-.5 1-1V5c0-.5.5-1 1-1s1 .5 1 1v12c0 .5.5 1 1 1s1-.5 1-1V4h1c.5 0 1-.5 1-1s-.5-1-1-1zm1 4v8l5-4l-5-4z"/>`),
		g.Group(children),
	)
}