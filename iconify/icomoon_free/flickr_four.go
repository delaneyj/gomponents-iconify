package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlickrFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0C3.582 0 0 3.606 0 8.055s3.582 8.055 8 8.055s8-3.606 8-8.055S12.418 0 8 0zM4.5 10.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5zm7 0a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z"/>`),
		g.Group(children),
	)
}