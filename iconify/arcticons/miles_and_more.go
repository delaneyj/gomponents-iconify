package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MilesAndMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2h0v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2h0v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 28.133v-6.696l3.352 6.703l3.351-6.693v6.693"/><circle cx="18.012" cy="21.646" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.012 23.699v4.441m9.268-.374a1.885 1.885 0 0 0 1.379.375h.376a1.11 1.11 0 0 0 1.108-1.11h0a1.109 1.109 0 0 0-1.108-1.11h-.752a1.11 1.11 0 0 1-1.108-1.11h0a1.11 1.11 0 0 1 1.108-1.11h.376a1.885 1.885 0 0 1 1.378.374m-4.695 3.22a1.675 1.675 0 0 1-1.456.845h0a1.676 1.676 0 0 1-1.676-1.675v-1.09a1.676 1.676 0 0 1 1.676-1.676h0a1.676 1.676 0 0 1 1.676 1.676v.545H22.21m-2.388-4.483v5.865a.838.838 0 0 0 .837.839h.252m-2.899 8.352v-6.696l3.352 6.703l3.352-6.693V36.5"/><rect width="3.352" height="4.441" x="26.415" y="32.059" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.676"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.511 33.735a1.676 1.676 0 0 1 1.676-1.676h0m-1.676 0V36.5m6.037-.846a1.675 1.675 0 0 1-1.456.846h0a1.676 1.676 0 0 1-1.676-1.676v-1.09a1.676 1.676 0 0 1 1.676-1.675h0a1.676 1.676 0 0 1 1.676 1.676v.545h-3.352M15.244 36.5h-.567a1.173 1.173 0 0 1-.948-.483l-2.325-3.196a2.745 2.745 0 0 1-.825-1.674a1.366 1.366 0 0 1 1.403-1.35a1.34 1.34 0 0 1 1.359 1.35c0 .765-.62 1.42-1.766 1.636c-1.277.243-2.075.943-2.075 2.098a1.563 1.563 0 0 0 1.766 1.619c1.547 0 2.457-1.467 3.802-3.27"/>`),
		g.Group(children),
	)
}