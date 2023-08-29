package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkPolicyOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="7" cy="2" r="1" fill="currentColor"/><circle cx="3" cy="6" r="1" fill="currentColor"/><circle cx="12" cy="1" r="1" fill="currentColor"/><circle cx="17" cy="2" r="1" fill="currentColor"/><circle cx="21" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="m17.5 24l-.119-.029A8.623 8.623 0 0 1 11 15.916v-4.218L17.5 9l6.5 2.698v4.218a8.623 8.623 0 0 1-6.381 8.055ZM13 12.865v3.15a6.396 6.396 0 0 0 4.5 5.96a6.396 6.396 0 0 0 4.5-5.96v-3.15l-4.5-1.793Zm10-.503Z"/><circle cx="12" cy="23" r="1" fill="currentColor"/><path fill="currentColor" d="M10.4 10.4a.802.802 0 0 0 .8-.8V8h1.6a1.6 1.6 0 0 0 1.59-1.5a6.027 6.027 0 0 1 2.353 1.857L18 7.833l1.08.45a7.997 7.997 0 1 0-8.39 11.6a9.57 9.57 0 0 1-.59-2.201a5.961 5.961 0 0 1-3.995-6.777L9.6 14.4v.8a1.585 1.585 0 0 0 .4 1.045V12H8.8v-1.6Z"/><circle cx="3" cy="18" r="1" fill="currentColor"/><circle cx="7" cy="22" r="1" fill="currentColor"/><circle cx="1" cy="12" r="1" fill="currentColor"/>`),
		g.Group(children),
	)
}