package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Callsblacklist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.69 0 15.25-7.77 15.25-16.94v-20c-4 0-15.25-2-15.25-2s-11.26 2-15.25 2v20C8.75 35.73 22.31 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.12 28.91a30.89 30.89 0 0 1 .54-3.05a1.52 1.52 0 0 0-.28-1.46c-.26-.25-1-1-1.74-1.76A14.06 14.06 0 0 1 22.27 19A13.81 13.81 0 0 1 26 16.32c.76.74 1.5 1.47 1.75 1.73a1.52 1.52 0 0 0 1.46.28a29 29 0 0 1 3-.53a1 1 0 0 0 .86-1.12v-4.1h0a.74.74 0 0 0-.68-.66c-3.71-.27-7.69 1.81-8.44 2.23h0l-.1.06h0a17.12 17.12 0 0 0-6.27 6.29h0l-.11.21h0c-.51.93-2.45 4.77-2.18 8.36a.72.72 0 0 0 .65.67H20a1 1 0 0 0 1.12-.83Zm11.97.86L15.22 11.9"/>`),
		g.Group(children),
	)
}