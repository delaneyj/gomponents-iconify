package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><g transform="translate(3 3)"><path d="M7.437 1.077L8.954-.441l.87.87l-1.518 1.517zM.883 10.403l-.868-.87l1.518-1.516l.868.87z"/><ellipse cx="5.115" cy="5.141" rx="3.115" ry="3.141"/><path fill="currentColor" d="M4.988 3.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977z"/></g><path d="M3.593 11.137h1.229v2.146H3.593z"/><ellipse cx="8.115" cy="8.142" rx="3.115" ry="3.142"/><path d="M11.114 2.744h1.229V4.89h-1.229z"/><path fill="currentColor" d="m11.991 10.911l2.384 1.845a8.022 8.022 0 0 0 1.583-4.787c0-4.413-3.573-7.906-7.979-7.906C3.573.063 0 3.556 0 7.969c0 4.412 3.572 7.989 7.979 7.989a7.976 7.976 0 0 0 4.013-1.068v-3.979h-.001zm.27-8.02l.869.869l-1.518 1.518l-.869-.869l1.518-1.518zM3.633 13.193l-.868-.869l1.518-1.518l.868.869l-1.518 1.518zM8 11.125c-1.721 0-3.115-1.406-3.115-3.14C4.885 6.25 6.28 4.843 8 4.843c1.719 0 3.115 1.406 3.115 3.142c0 1.734-1.396 3.14-3.115 3.14z"/><path fill="currentColor" d="M13.031 16v-2.969l1.938 1.5L13.031 16z"/></g>`),
		g.Group(children),
	)
}