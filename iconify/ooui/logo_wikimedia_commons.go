package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWikimediaCommons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.09 6.18a3.68 3.68 0 0 1-2.18-2.55c.09.09 1.82.91 1.82.91L10 0L7.27 4.55l1.82-.91a5.08 5.08 0 0 0 .55 1.91a5.13 5.13 0 0 0 2 2a8.86 8.86 0 0 1 2 1.18l-.64.63l-.45-.45l-.26 1.54l1.54-.26l-.45-.45l.62-.65a5.69 5.69 0 0 1 1.45 3.45h-.91v-.73l-1.26.91l1.26.91v-.73h.91A5.21 5.21 0 0 1 14 16.36l-.64-.64l.45-.45l-1.53-.27l.26 1.54l.45-.45l.64.64a5.69 5.69 0 0 1-3.45 1.45v-.91h.73L10 16l-.91 1.27h.73v.91a5.21 5.21 0 0 1-3.45-1.45l.63-.64l.45.45l.27-1.54l-1.54.26l.45.45l-.63.65a5.69 5.69 0 0 1-1.45-3.45h.91v.73l1.26-.91l-1.26-.91v.73h-.91A5.21 5.21 0 0 1 6 9.09l.64.64l-.45.45l1.54.26l-.28-1.53l-.45.45l-.64-.64L5 7.45a7.29 7.29 0 1 0 8.09-1.27z"/><circle cx="10" cy="12.7" r="2.5" fill="currentColor"/>`),
		g.Group(children),
	)
}