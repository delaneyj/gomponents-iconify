package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Degrees(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M13.19 9.21c0-.5.18-.93.53-1.28c.36-.36.78-.53 1.28-.53c.49 0 .92.18 1.27.53c.35.36.53.78.53 1.28s-.18.93-.53 1.29c-.35.36-.78.54-1.27.54s-.92-.18-1.28-.54s-.53-.79-.53-1.29zm.88 0c0 .26.09.48.27.67c.19.19.41.28.67.28c.26 0 .48-.09.67-.28s.28-.41.28-.67a.87.87 0 0 0-.28-.66a.947.947 0 0 0-.67-.28c-.26 0-.48.09-.67.27c-.18.18-.27.4-.27.67z"/>`),
		g.Group(children),
	)
}