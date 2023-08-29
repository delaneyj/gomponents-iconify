package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterLowercaseCircleG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M29.56 12.879V29.56A5.56 5.56 0 0 1 24 35.121h0a5.543 5.543 0 0 1-3.932-1.629"/><rect width="11.121" height="14.735" x="18.439" y="12.879" rx="5.561" ry="5.561" transform="rotate(180 24 20.247)"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}