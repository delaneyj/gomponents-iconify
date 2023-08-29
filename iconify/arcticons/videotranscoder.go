package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Videotranscoder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.78 32a8.65 8.65 0 0 0 .08-1.15a8.77 8.77 0 0 0-.08-1.15l2.5-2a.57.57 0 0 0 .14-.75l-2.36-4.1a.59.59 0 0 0-.72-.25l-2.94 1.17a9.24 9.24 0 0 0-2-1.16l-.4-3.13a.6.6 0 0 0-.59-.5h-4.78a.6.6 0 0 0-.58.5l-.44 3.13a8.85 8.85 0 0 0-2 1.16l-2.95-1.18a.58.58 0 0 0-.72.25l-2.36 4.1a.59.59 0 0 0 .14.75l2.5 2a8.91 8.91 0 0 0-.09 1.15A8.64 8.64 0 0 0 7.22 32l-2.5 2a.6.6 0 0 0-.14.76l2.36 4.09a.59.59 0 0 0 .72.26l2.95-1.19a8.64 8.64 0 0 0 2 1.16l.44 3.14a.6.6 0 0 0 .58.49h4.73a.6.6 0 0 0 .59-.49l.44-3.22a8.49 8.49 0 0 0 2-1.16L24.33 39a.59.59 0 0 0 .72-.26l2.37-4.09a.6.6 0 0 0-.14-.76ZM16 35a4.14 4.14 0 1 1 4.14-4.14A4.14 4.14 0 0 1 16 34.94Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.69 33.45H40.2a3.3 3.3 0 0 0 3.3-3.3V5.38h-7.14l3.94 5.74h-4.91l-3.94-5.74H27.5l3.95 5.74h-4.91l-3.95-5.74h-3.94l3.94 5.74h-4.91l-3.94-5.74H11.5a3.3 3.3 0 0 0-3.3 3.3V22.8"/>`),
		g.Group(children),
	)
}