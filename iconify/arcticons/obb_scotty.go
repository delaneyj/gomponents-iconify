package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObbScotty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.933 32.077h2.541v2.371h-2.541zm14.026 0H43.5v2.371h-2.541zm-5.742-18.525c-2.953 0-8.284.459-8.284 1.66v16.865H43.5V15.211c0-1.2-5.37-1.659-8.283-1.659Zm.012 2.778a22.418 22.418 0 0 1 6.087.61a.48.48 0 0 1 .333.463v8.493a.487.487 0 0 1-.487.486H29.296a.487.487 0 0 1-.486-.486v-8.493a.48.48 0 0 1 .332-.462a22.422 22.422 0 0 1 6.087-.611Zm-5.207 11.363a1.248 1.248 0 0 1 0 2.496h0a1.23 1.23 0 0 1-1.212-1.248h0a1.23 1.23 0 0 1 1.212-1.248Zm10.415 0a1.248 1.248 0 0 1 0 2.496h0a1.23 1.23 0 0 1-1.212-1.248h0a1.23 1.23 0 0 1 1.212-1.248Z"/><circle cx="9.144" cy="27.187" r="1.324" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="18.332" cy="27.187" r="1.324" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.617 16.433v12.119a2.887 2.887 0 0 0 2.88 2.88H18.98a2.887 2.887 0 0 0 2.88-2.88v-12.12a2.887 2.887 0 0 0-2.88-2.88H8.498a2.887 2.887 0 0 0-2.88 2.88ZM8.498 31.46L4.5 34.448m14.479-2.988l3.997 2.988"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.776 16.248H19.7v6.276H7.776z"/>`),
		g.Group(children),
	)
}