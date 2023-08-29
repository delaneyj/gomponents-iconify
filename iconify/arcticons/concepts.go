package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Concepts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="37.008" cy="8.308" r="3.599" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.03" d="M16.337 42.804A20.393 20.393 0 1 1 34.09 6.201"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.4 31.834l1.222.794m-4.018 4.012l1.142.985m-5.436 2.757l.788 1.025m-6.077 1.313l.404 1.253m-5.712-.636l-.04 1.194m-5.515-1.791l-.275 1.188M45.5 27.264h-5.438l-4.964-1.497l.079-4.097l4.885-1.497H45.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.098 25.767l-5.515-2.127l5.594-1.97"/>`),
		g.Group(children),
	)
}