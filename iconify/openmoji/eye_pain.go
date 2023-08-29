package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyePain(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36.144" cy="39.651" r="8.896" fill="#a57939"/><ellipse cx="36" cy="39.921" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="20" ry="11.737"/><circle cx="36.144" cy="39.651" r="4.85"/><circle cx="36.144" cy="39.651" r="4.85" fill="none"/><circle cx="36.144" cy="39.651" r="8.896" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m43.135 24.35l1.983-4.037l2.026 2.152l2.491-5.072m-25.319-1.134l2.333 3.845l-2.904.549l2.93 4.83m9.805-2.801l-.19-4.494l-2.721 1.154l-.238-5.646"/><path fill="none" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m43.135 24.35l1.983-4.037l2.026 2.152l2.491-5.072m-25.319-1.134l2.333 3.845l-2.904.549l2.93 4.83m9.805-2.801l-.19-4.494l-2.721 1.154l-.238-5.646"/>`),
		g.Group(children),
	)
}