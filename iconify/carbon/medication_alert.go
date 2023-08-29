package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicationAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 19h2v5h-2z"/><path fill="currentColor" d="M23 31a8 8 0 1 1 8-8a8 8 0 0 1-8 8Zm0-14a6 6 0 1 0 6 6a6 6 0 0 0-6-6Z"/><circle cx="23" cy="26" r="1" fill="currentColor"/><path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2v18a2 2 0 0 0 2 2h8v-2H6v-2h5V12H6v-2h12v3h2v-3a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM9 14v10H6V14ZM4 8V4h16v4Z"/>`),
		g.Group(children),
	)
}