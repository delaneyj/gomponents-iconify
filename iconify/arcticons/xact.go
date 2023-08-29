package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xact(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.032 21.509l2.592 2.497l-2.592 2.497a7.115 7.115 0 0 1-9.91 0l-2.678-2.41l-10.627 9.909c-1.693 1.607-1.788 4.285-.086 5.892c1.607 1.694 4.285 1.694 5.979.087l12.354-11.387l12.294 11.395c1.607 1.694 4.286 1.607 5.893 0c.803-.803 1.157-1.875 1.157-2.86c0-1.07-.449-2.228-1.339-3.032L31.624 24.006l-2.592-2.497"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.48 15.876c0-1.071.355-2.056 1.158-2.86c1.607-1.693 4.286-1.52 5.979.086l7.007 6.42c-1.694 0-3.344.812-4.683 2.065l-2.678 2.497l-5.538-5.09c-.795-.88-1.244-2.047-1.244-3.118Zm15.716 3.732a7.548 7.548 0 0 0-.985-.069L36.36 7.936c1.693-1.607 4.371-1.52 5.978.086c1.607 1.694 1.52 4.286-.086 5.893l-10.627 10.09h0l-2.592-2.496h0s-.717-.7-2.047-1.322c-1.322-.622-1.789-.579-1.789-.579Z"/>`),
		g.Group(children),
	)
}