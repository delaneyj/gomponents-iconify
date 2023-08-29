package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCalculatorOne0"><g fill="none"><rect width="32" height="40" x="8" y="4" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 11h20v9H14z"/><circle cx="17" cy="26" r="2" fill="#fff"/><circle cx="17" cy="32" r="2" fill="#fff"/><circle cx="17" cy="38" r="2" fill="#fff"/><circle cx="24" cy="26" r="2" fill="#fff"/><circle cx="24" cy="32" r="2" fill="#fff"/><circle cx="24" cy="38" r="2" fill="#fff"/><circle cx="31" cy="26" r="2" fill="#fff"/><circle cx="31" cy="32" r="2" fill="#fff"/><circle cx="31" cy="38" r="2" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCalculatorOne0)"/>`),
		g.Group(children),
	)
}