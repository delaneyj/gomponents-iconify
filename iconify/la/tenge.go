package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tenge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6.633 5A.625.625 0 0 0 6 5.621v2.746c0 .348.285.633.633.633h18.746A.625.625 0 0 0 26 8.367V5.621A.616.616 0 0 0 25.379 5H6.633zm0 7a.632.632 0 0 0-.633.621v2.746c0 .348.285.633.633.633H14v11.379c0 .347.274.621.621.621h2.758a.616.616 0 0 0 .621-.621V16h7.379a.625.625 0 0 0 .621-.633v-2.746a.622.622 0 0 0-.621-.621H6.633z"/>`),
		g.Group(children),
	)
}