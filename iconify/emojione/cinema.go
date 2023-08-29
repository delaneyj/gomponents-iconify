package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M42 16c-3 0-5.6 1.7-7 4.1c-1.4-2.5-4-4.1-7-4.1c-4.4 0-8 3.6-8 8s3.6 8 8 8c3 0 5.6-1.7 7-4.1c1.4 2.5 4 4.1 7 4.1c4.4 0 8-3.6 8-8s-3.6-8-8-8M28 27c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3m14 0c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3m4 5H23c-1.1 0-2 .9-2 2v.5L16 32v16l5-2.5v.5c0 1.1.9 2 2 2h23c1.1 0 2-.9 2-2V34c0-1.1-.9-2-2-2"/>`),
		g.Group(children),
	)
}