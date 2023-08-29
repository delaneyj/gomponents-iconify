package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yadrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 8.32c.89-.1 2.28-.25 3.1-.28c5.1-.2 8.82.81 11.38 2.72c3.43 2.54 4.6 6.62 4.24 10.92c-.64 7.62-5.78 16.32-12.1 20.82m12.16 0c7.9-6.85 12.33-14 13.39-20.53c1-6.22-1.06-11.9-5.65-16.47"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.01"/>`),
		g.Group(children),
	)
}