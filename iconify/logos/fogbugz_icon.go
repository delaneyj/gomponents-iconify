package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FogbugzIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosFogbugzIcon0" x1="-24.513%" x2="130.064%" y1="75.564%" y2="-56.513%"><stop offset="0%" stop-color="#AB68FC"/><stop offset="86%" stop-color="#5521B3"/></linearGradient><linearGradient id="logosFogbugzIcon1" x1="104.596%" x2="48.252%" y1="100.586%" y2="48.39%"><stop offset="2%" stop-color="#AA69FF"/><stop offset="86%" stop-color="#C6B1FF"/></linearGradient></defs><path fill="url(#logosFogbugzIcon0)" d="M89.976 306.862L256 161.802H90.258s-.246 139.468-.282 145.06Z"/><path fill="url(#logosFogbugzIcon1)" d="M166.059 0L0 161.803h165.32S166.024 4.22 166.06 0Z"/>`),
		g.Group(children),
	)
}