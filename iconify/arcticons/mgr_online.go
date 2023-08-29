package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MgrOnline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.835 31.979h0a1.972 1.972 0 0 1-1.986-1.958v-2.092c0-1.145.917-2.062 1.986-2.062h0a2.055 2.055 0 0 1 2.063 2.046v2.003a2.055 2.055 0 0 1-2.047 2.063h-.016Zm26.316 0h-3.056v-6.112h3.056m-3.056 3.056h1.986m-8.044 3.056v-6.112l4.049 6.112v-6.112m-17.18 6.112v-6.112l4.05 6.112v-6.112m7.072 0v6.112m-5.064-6.112v6.112h3.056m6.209-15.949h5.896c.762 0 1.374.656 1.374 1.503v.61c0 .846-.612 1.528-1.375 1.528H31.23m5.894.026c.763 0 1.376.656 1.376 1.503v1.27m-7.27-6.441v6.441m-1.978-5.136c-.131-.744-.926-1.313-1.892-1.313h-3.957c-1.06 0-1.91.681-1.91 1.528v3.401c0 .847.85 1.528 1.91 1.528h3.957c1.058 0 1.91-.68 1.91-1.528v-1.24h-3.882m-5.855 2.746v-6.413h-.541l-3.92 6.413h-1.111l-3.92-6.413H9.5v6.413"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}