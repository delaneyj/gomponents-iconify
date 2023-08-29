package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPortugal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 17h21v38H5z"/><circle cx="26" cy="36" r="12" fill="none" stroke="#fcea2b" stroke-miterlimit="10" stroke-width="2"/><path fill="none" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="M26 24v24m0-8.5L17 44h18l-9-4.5zm0-6l9-5.5l-8.5 1.5L17 28l9 5.5z"/><path fill="none" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="m38 36l-12 5l-12-5l12-5l12 5z"/><path fill="#fff" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.2 29h11.6v10.1c0 2.5-2.6 4.6-5.8 4.6h0c-3.2 0-5.8-2.1-5.8-4.6V29Z"/><circle cx="26" cy="32.8" r=".7" fill="#1e50a0" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26" cy="38.7" r=".7" fill="#1e50a0" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26" cy="35.7" r=".7" fill="#1e50a0" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29" cy="35.7" r=".7" fill="#1e50a0" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round"/><circle cx="23" cy="35.7" r=".7" fill="#1e50a0" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}