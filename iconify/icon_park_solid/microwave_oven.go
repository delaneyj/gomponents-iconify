package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrowaveOven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMicrowaveOven0"><g fill="none"><rect width="40" height="30" x="4" y="6" stroke="#fff" stroke-width="4" rx="2"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 15h-4m4 6h-4"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 15h17v12H10z"/><circle cx="36" cy="27" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 36v6m24-6v6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMicrowaveOven0)"/>`),
		g.Group(children),
	)
}