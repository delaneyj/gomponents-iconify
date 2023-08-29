package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscPause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><g transform="translate(3 3)"><path d="M7.466 1.064L8.984-.453l.869.869l-1.518 1.517zm-6.57 9.341l-.869-.87L1.546 8.02l.868.87z"/><ellipse cx="5.115" cy="5.141" rx="3.115" ry="3.141"/><path fill="currentColor" d="M4.988 3.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977z"/></g><path fill="currentColor" d="M15.301 10.953c.4-.991.657-1.967.657-2.984c0-4.413-3.573-7.906-7.979-7.906C3.573.063 0 3.556 0 7.969c0 4.412 3.572 7.989 7.979 7.989c.68 0 1.332-.1 1.959-.264v-4.741h5.363zm-3.04-8.062l.869.869l-1.518 1.518l-.869-.869l1.518-1.518zM3.633 13.193l-.868-.869l1.518-1.518l.868.869l-1.518 1.518zM8 5a3 3 0 1 1-.002 6.002A3 3 0 0 1 8 5z"/><path fill="currentColor" d="M11 12h2.009v4H11zm3 0h1.981v4H14z"/></g>`),
		g.Group(children),
	)
}