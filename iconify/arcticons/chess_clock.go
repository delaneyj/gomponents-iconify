package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessClock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.226 36.142c.6.5 1.1.7 2.4.7h.3c1.4 0 2.6-1.2 2.6-2.6s-1.2-2.6-2.6-2.6h-2.7v-2.8h5.3"/><circle cx="20.142" cy="35.639" r=".75" fill="currentColor"/><circle cx="20.142" cy="30.339" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.758 34.369c0 1.5 1.2 2.7 2.7 2.7s2.6-1.2 2.6-2.7v-2.7c0-1.5-1.2-2.7-2.6-2.7s-2.7 1.2-2.7 2.7v2.7Zm7.716 0c0 1.5 1.2 2.7 2.7 2.7s2.6-1.2 2.6-2.7v-2.7c0-1.5-1.2-2.7-2.6-2.7s-2.7 1.2-2.7 2.7v2.7Z" class="c"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24h37m-6.726-12.142c-.6-.5-1.1-.7-2.4-.7h-.3c-1.4 0-2.6 1.2-2.6 2.6s1.2 2.6 2.6 2.6h2.7v2.8h-5.3"/><circle cx="27.858" cy="12.361" r=".75" fill="currentColor"/><circle cx="27.858" cy="17.661" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.242 13.63c0-1.5-1.2-2.7-2.7-2.7s-2.6 1.2-2.6 2.7v2.7c0 1.5 1.2 2.7 2.6 2.7s2.7-1.2 2.7-2.7v-2.7Zm-7.716 0c0-1.5-1.2-2.7-2.7-2.7s-2.6 1.2-2.6 2.7v2.7c0 1.5 1.2 2.7 2.6 2.7s2.7-1.2 2.7-2.7v-2.7Z" class="c"/>`),
		g.Group(children),
	)
}