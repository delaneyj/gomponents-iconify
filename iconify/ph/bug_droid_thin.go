package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BugDroidThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m186.17 43.49l16.66-16.66a4 4 0 1 0-5.66-5.66l-17 17a83.72 83.72 0 0 0-104.26 0l-17-17a4 4 0 0 0-5.66 5.66l16.58 16.66A83.75 83.75 0 0 0 44 104v40a84 84 0 0 0 168 0v-40a83.75 83.75 0 0 0-25.83-60.51ZM128 28a76.08 76.08 0 0 1 76 76v12H52v-12a76.08 76.08 0 0 1 76-76Zm0 192a76.08 76.08 0 0 1-76-76v-20h152v20a76.08 76.08 0 0 1-76 76Zm20-136a8 8 0 1 1 8 8a8 8 0 0 1-8-8Zm-56 0a8 8 0 1 1 8 8a8 8 0 0 1-8-8Z"/>`),
		g.Group(children),
	)
}