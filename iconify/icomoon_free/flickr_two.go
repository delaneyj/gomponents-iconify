package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlickrTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.5 6.5c-1.103 0-2 .897-2 2s.897 2 2 2s2-.897 2-2s-.897-2-2-2zm0-1.5a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7zM0 8.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0z"/>`),
		g.Group(children),
	)
}