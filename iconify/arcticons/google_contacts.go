package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleContacts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.72 23.085h4.433c4.826 0 8.743 3.918 8.743 8.743v2.929c0 4.825-3.917 8.743-8.743 8.743H9.863a2.238 2.238 0 0 1-2.238-2.237V36.18c0-7.227 5.868-13.095 13.095-13.095Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.847 23.085h8.784c4.826 0 8.744 3.918 8.744 8.743v2.929c0 4.825-3.918 8.743-8.744 8.743h-15.29a2.238 2.238 0 0 1-2.237-2.237V31.83c0-4.826 3.918-8.744 8.743-8.744Z"/><circle cx="23.892" cy="12.186" r="7.686" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}