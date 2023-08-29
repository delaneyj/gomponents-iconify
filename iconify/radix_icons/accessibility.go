package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accessibility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M.877 7.5a6.623 6.623 0 1 1 13.246 0a6.623 6.623 0 0 1-13.246 0ZM7.5 1.827a5.673 5.673 0 1 0 0 11.346a5.673 5.673 0 0 0 0-11.346ZM7.125 9c-.055.127-.793 2.96-.793 2.96a.5.5 0 1 1-.966-.26s.88-2.827.88-3.43V6.801l-1.958-.525a.5.5 0 1 1 .258-.966s1.654.563 2.3.563h1.309c.645 0 2.298-.563 2.298-.563a.5.5 0 1 1 .26.966l-1.966.527V8.27c0 .603.88 3.427.88 3.427a.5.5 0 0 1-.966.259S7.92 9.127 7.869 9c-.05-.127-.219-.127-.219-.127h-.307s-.173 0-.218.127ZM7.5 5.12a1.125 1.125 0 1 0 0-2.25a1.125 1.125 0 0 0 0 2.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}