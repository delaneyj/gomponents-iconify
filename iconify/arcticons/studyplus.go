package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Studyplus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.473 5.357l2.517 2.518l-11.971 11.97L5.5 17.33zm21.411 15.71l2.517 2.517l-14.828 14.828l-2.517-2.517zm-15.74-2.823l2.517 2.517l-8.365 8.365l-2.517-2.517zm7.966 1.303l2.517 2.518L21.932 33.76l-2.517-2.517zm-11.644-6.888l2.518 2.517l-9.32 9.32l-2.517-2.518zM40.66 28.56l1.84 14.08l-13.81-2.11l11.97-11.97zm-2.74 13.36l3.86-3.87"/>`),
		g.Group(children),
	)
}