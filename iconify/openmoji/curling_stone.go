package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurlingStone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M30 17h16v3H30zm-2 0h2v6h-4v-4a2 2 0 0 1 2-2Zm-4 8h25v3H24z"/><path fill="#9b9b9a" d="m15.024 45.316l43.782-.014A5.12 5.12 0 0 1 59 47.931a6.665 6.665 0 0 1-3 3.899a12.842 12.842 0 0 1-4.46 2.11c-6.418 1.59-14.579.957-22.54.814a196.63 196.63 0 0 0-7 0a11.766 11.766 0 0 1-6-2.924c-1.958-1.84-3.823-5.016-2.899-6.127c.414-.497 1.297-.47 1.922-.387ZM59 40H13a5.281 5.281 0 0 1-1-3.385c.224-2.978 3.408-5.721 8-6.771c4.221-.455 8.91-.778 14-.846a149.335 149.335 0 0 1 18 .846c4.98 1.589 8.173 4.656 8 7.617A4.665 4.665 0 0 1 59 40Z"/><path fill="#3f3f3f" d="M47.154 28A14.055 14.055 0 0 0 41 29.42a13.332 13.332 0 0 1 0 24.16A14.055 14.055 0 0 0 47.154 55a13.504 13.504 0 1 0 0-27Z"/><path fill="#9b9b9a" d="M60 42.5a14.324 14.324 0 0 0-.227-2.5H48.446a13.898 13.898 0 0 1 0 5h11.327A14.325 14.325 0 0 0 60 42.5Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M26 17h4v7h-4z"/><rect width="48" height="25.6" x="12" y="29.067" rx="12"/><rect width="48" height="5" x="12" y="40" rx="2.5"/><rect width="26" height="4" x="23" y="25" rx="2"/><path d="M44.5 20H30v-3h14.5a1.5 1.5 0 0 1 0 3Z"/></g>`),
		g.Group(children),
	)
}