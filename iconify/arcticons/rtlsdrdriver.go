package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rtlsdrdriver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.79 6.35a.9.9 0 0 0-.9.91h0v1.62a.47.47 0 0 1-.36.44a11.62 11.62 0 0 0-4.19 1.74a.42.42 0 0 1-.56-.06l-1.15-1.13a.9.9 0 0 0-1.27 0l-1.71 1.72a.91.91 0 0 0 0 1.28L13.78 14a.43.43 0 0 1 .05.57a12.19 12.19 0 0 0-1.73 4.23a.44.44 0 0 1-.43.36h-1.6a.91.91 0 0 0-.91.91h0v2.43a.91.91 0 0 0 .9.92h1.61a.44.44 0 0 1 .43.36A12.19 12.19 0 0 0 13.83 28a.43.43 0 0 1-.05.57l-1.13 1.17a.91.91 0 0 0 0 1.28l1.71 1.72a.9.9 0 0 0 1.27 0l1.13-1.13a.43.43 0 0 1 .57-.06a11.65 11.65 0 0 0 4.18 1.74a.46.46 0 0 1 .36.44v1.61a.9.9 0 0 0 .9.92h2.42a.91.91 0 0 0 .91-.91h0v-1.61a.45.45 0 0 1 .36-.44a11.77 11.77 0 0 0 4.19-1.74a.42.42 0 0 1 .56.06l1.16 1.13a.9.9 0 0 0 1.27 0L35.35 31a.91.91 0 0 0 0-1.28l-1.13-1.14a.43.43 0 0 1 0-.57a12.19 12.19 0 0 0 1.73-4.22a.44.44 0 0 1 .43-.36h1.6a.91.91 0 0 0 .91-.91h0v-2.43a.9.9 0 0 0-.9-.91h-1.66a.44.44 0 0 1-.43-.36a12.11 12.11 0 0 0-1.73-4.22a.43.43 0 0 1 0-.57l1.13-1.16a.91.91 0 0 0 0-1.28l-1.66-1.72a.9.9 0 0 0-1.27 0L31.24 11a.43.43 0 0 1-.57.06a11.65 11.65 0 0 0-4.18-1.74a.45.45 0 0 1-.36-.44V7.27a.9.9 0 0 0-.9-.92h-2.44ZM24 14.94h0a6.37 6.37 0 1 1-6.31 6.37A6.34 6.34 0 0 1 24 14.94Z"/><ellipse cx="24" cy="21.31" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2.02"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.19 22.12l-1 4.77m-1.28 6l-2.3 10.61h1L24 39.19l5.36 4.31h1l-2.25-10.64m-1.28-6L26 23"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.66 35.44L24 39.19l-4.66-3.75M24 23.7l-2.28 1.83L24 27.36l2.28-1.83Z"/><ellipse cx="24" cy="21.31" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="16.67" ry="16.81"/>`),
		g.Group(children),
	)
}