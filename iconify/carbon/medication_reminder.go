package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicationReminder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2v18a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V10a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM6 14h3v10H6Zm12 14H6v-2h5V12H6v-2h12ZM4 8V4h16v4Z"/><circle cx="26" cy="16" r="4" fill="currentColor"/>`),
		g.Group(children),
	)
}