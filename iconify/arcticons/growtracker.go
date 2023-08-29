package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Growtracker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.07 26.41h-.19v.19a16.93 16.93 0 0 0 16.94 16.9H24v-.18A16.92 16.92 0 0 0 7.07 26.41ZM41 26.4a16.92 16.92 0 0 0-16.9 16.92v.18h.18a16.92 16.92 0 0 0 16.9-16.93v-.17ZM28.64 26a4 4 0 0 0 0-.49a4.4 4.4 0 0 0 1 .65a4.65 4.65 0 1 0 4-8.39l-.44-.2l.44-.2a5.46 5.46 0 0 0 .59-.39a4.62 4.62 0 1 0-5.59-7.36a4 4 0 0 0 0-.49a4.64 4.64 0 0 0-9.27 0a4 4 0 0 0 0 .49a4.77 4.77 0 0 0-1-.65a4.65 4.65 0 1 0-4 8.4l.44.2l-.44.2a5.36 5.36 0 0 0-.59.38a4.62 4.62 0 1 0 5.59 7.36a4 4 0 0 0 0 .49a4.64 4.64 0 0 0 9.27 0ZM24 22.19a4.64 4.64 0 1 1 4.64-4.64A4.63 4.63 0 0 1 24 22.19Z"/>`),
		g.Group(children),
	)
}