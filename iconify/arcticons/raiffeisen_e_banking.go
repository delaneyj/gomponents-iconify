package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RaiffeisenEBanking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="25.942" cy="29.055" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="10.407" ry="4.337"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.349 31.837v-2.782m-20.815 0v4.337c0 2.395 4.66 4.337 10.408 4.337a21.935 21.935 0 0 0 5.716-.712M22.687 16.85a19.076 19.076 0 0 0-7.78-1.456C9.16 15.394 4.5 17.335 4.5 19.73s4.66 4.337 10.407 4.337a19.078 19.078 0 0 0 7.78-1.456"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 19.73v4.338c0 2.395 4.66 4.337 10.407 4.337q.384 0 .761-.011"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 24.068v4.337c0 2.396 4.66 4.338 10.407 4.338q.316 0 .627-.008"/><ellipse cx="33.093" cy="14.608" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="10.407" ry="4.337"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.685 14.608v4.337c0 2.396 4.66 4.338 10.408 4.338S43.5 21.34 43.5 18.945v-4.337"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.642 27.489c4.515-.474 7.858-2.177 7.858-4.206v-4.338m-20.815 0v4.338a2.311 2.311 0 0 0 .704 1.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.35 31.74c4.152-.569 7.15-2.199 7.15-4.12v-4.337"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.85 42.5v-8h2.619a2.687 2.687 0 0 1 0 5.373H35.85m2.619 0l2.619 2.625"/>`),
		g.Group(children),
	)
}