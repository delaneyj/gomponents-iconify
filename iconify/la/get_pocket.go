package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GetPocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 5C5.355 5 4 6.355 4 8v7c0 6.617 5.383 12 12 12s12-5.383 12-12V8c0-1.645-1.355-3-3-3zm0 2h18c.566 0 1 .434 1 1v7c0 5.535-4.465 10-10 10S6 20.535 6 15V8c0-.566.434-1 1-1zm3.656 4.406c-.383 0-.77.176-1.062.469a1.467 1.467 0 0 0 0 2.094L15 19.375a1.504 1.504 0 0 0 2.125 0l5.281-5.25a1.504 1.504 0 0 0 0-2.125a1.504 1.504 0 0 0-2.125 0l-4.218 4.219l-4.344-4.344c-.293-.293-.68-.469-1.063-.469z"/>`),
		g.Group(children),
	)
}