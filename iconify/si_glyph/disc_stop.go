package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><g transform="translate(3 3)"><path d="M7.465 1.063L8.982-.455l.87.87l-1.518 1.517zm-6.57 9.34l-.868-.87l1.518-1.516l.869.87z"/><ellipse cx="5.115" cy="5.141" rx="3.115" ry="3.141"/><path fill="currentColor" d="M4.988 3.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977z"/></g><path fill="currentColor" d="M15.301 10.953c.4-.991.657-1.967.657-2.984c0-4.413-3.573-7.906-7.979-7.906C3.573.063 0 3.556 0 7.969c0 4.412 3.572 7.989 7.979 7.989c.68 0 1.332-.1 1.959-.264l1.031-.257l-.031-4.484h4.363zm-3.04-8.062l.869.869l-1.518 1.518l-.869-.869l1.518-1.518zM3.633 13.193l-.868-.869l1.518-1.518l.868.869l-1.518 1.518zM5 8.001a3 3 0 1 1 6 0a3.001 3.001 0 0 1-6 0z"/><path fill="currentColor" d="M12 12h4.023v4H12z"/></g>`),
		g.Group(children),
	)
}