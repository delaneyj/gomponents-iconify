package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blackberry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1019.5 287.5q-12.5 39.5-49 68T894 384H735l64-192h154q40 0 59.5 28t7 67.5zM607 192H447L511 0h155q39 0 58.5 28t7.5 68t-48.5 68t-76.5 28zm-5 64q39 0 58.5 28t7.5 67.5t-48.5 67.5t-76.5 28H383l64-191h155zm-378-64H64L128 0h154q40 0 59.5 28t7.5 68t-49 68t-76 28zm-6 64q40 0 59.5 28t7.5 67.5t-49 67.5t-76 28H0l64-191h154zm320 255q39 0 58.5 28t7.5 68t-48.5 68t-76.5 28H319l64-192h155zm351-64q40 0 59.5 28.5t7.5 68t-49 67.5t-76 28H671l64-192h154z"/>`),
		g.Group(children),
	)
}