package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.47 3H3.53a.53.53 0 0 0-.53.53v16.94c0 .293.237.53.53.53h16.94a.53.53 0 0 0 .53-.53V3.53a.53.53 0 0 0-.53-.53zM11.5 19.941H4.059V12.5H11.5v7.441zm0-8.441H4.059V4.059H11.5V11.5zm8.441 8.441H12.5V12.5h7.441v7.441zm0-8.441H12.5V4.059h7.441V11.5z"/>`),
		g.Group(children),
	)
}