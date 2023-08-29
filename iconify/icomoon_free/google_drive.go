package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.844 10L3.96 15h9.072l2.884-5zm8.662-1l-4.619-8H5.112l4.619 8zM4.534 2L0 9.856l2.888 5L7.422 7z"/>`),
		g.Group(children),
	)
}