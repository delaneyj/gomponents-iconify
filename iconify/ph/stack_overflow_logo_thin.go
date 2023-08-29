package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackOverflowLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212 152.09V216a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4v-63.91a4 4 0 0 1 8 0V212h152v-59.91a4 4 0 0 1 8 0Zm-124 28h80a4 4 0 1 0 0-8H88a4 4 0 1 0 0 8Zm5.92-52.86l77.27 20.67a3.72 3.72 0 0 0 1 .14a4 4 0 0 0 1-7.85L96 119.47a4 4 0 1 0-2.07 7.72Zm19.41-49.54l69.28 39.95a4 4 0 1 0 4-6.92l-69.28-39.99a4 4 0 1 0-4 6.92Zm88.16 13.66a4 4 0 0 0 5.66-5.64l-56.57-56.5a4 4 0 1 0-5.66 5.65Z"/>`),
		g.Group(children),
	)
}