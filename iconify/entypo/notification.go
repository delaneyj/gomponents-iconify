package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 8.38V17H3V5h8.62c-.073-.322-.12-.655-.12-1s.047-.678.12-1H3c-1.102 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V8.38c-.322.073-.655.12-1 .12s-.678-.047-1-.12zM16 1a3 3 0 1 0 0 6a3 3 0 0 0 0-6z"/>`),
		g.Group(children),
	)
}