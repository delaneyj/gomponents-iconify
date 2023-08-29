package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm1 17.93V19a1 1 0 0 0-2 0v.93A8 8 0 0 1 4.07 13H5a1 1 0 0 0 0-2h-.93A8 8 0 0 1 11 4.07V5a1 1 0 0 0 2 0v-.93A8 8 0 0 1 19.93 11H19a1 1 0 0 0 0 2h.93A8 8 0 0 1 13 19.93Zm2.14-12.38l-5 2.12a1 1 0 0 0-.52.52l-2.12 5a1 1 0 0 0 .21 1.1a1 1 0 0 0 .7.3a.93.93 0 0 0 .4-.09l5-2.12a1 1 0 0 0 .52-.52l2.12-5a1 1 0 0 0-1.31-1.31Zm-2.49 5.1l-2.28 1l1-2.28l2.28-1Z"/>`),
		g.Group(children),
	)
}