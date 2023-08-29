package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 112q-53 0-90.5 37.5T43 240q0 32 15 60l-32 31Q0 289 0 240q0-71 50-121t121-50V5l85 86l-85 85v-64zm144 37q26 42 26 91q0 71-50 121t-120 50v64l-86-86l86-85v64q53 0 90.5-37.5T299 240q0-31-15-60z"/>`),
		g.Group(children),
	)
}