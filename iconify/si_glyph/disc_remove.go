package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M13 14h3.938v1.969H13z"/><path d="M16.958 7.969c0-4.413-3.573-7.906-7.979-7.906C4.573.063 1 3.556 1 7.969c0 4.412 3.572 7.989 7.979 7.989a7.881 7.881 0 0 0 2.992-.601v-2.402h3.314c.941-1.319 1.673-3.255 1.673-4.986zM4.633 13.193l-.868-.869l1.518-1.518l.868.869l-1.518 1.518zM9 11.125c-1.721 0-3.115-1.406-3.115-3.14C5.885 6.25 7.28 4.843 9 4.843c1.719 0 3.115 1.406 3.115 3.142c0 1.734-1.396 3.14-3.115 3.14zm3.612-5.848l-.869-.869l1.518-1.518l.869.869l-1.518 1.518z"/><path d="M8.988 6.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977z"/></g>`),
		g.Group(children),
	)
}