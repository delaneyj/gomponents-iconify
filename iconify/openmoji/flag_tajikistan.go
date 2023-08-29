package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTajikistan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 17h62v13H5z"/><path fill="#fff" d="M5 30h62v12H5z"/><circle cx="36" cy="35.155" r=".497" fill="#fcea2b" stroke="#fcea2b" stroke-miterlimit="10" stroke-width=".994"/><rect width=".928" height="3.975" x="35.536" y="35.155" fill="#fcea2b" rx=".306" ry=".306"/><path fill="none" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.988" d="M38.982 40.125a4.214 4.214 0 0 0-5.963 0"/><g fill="#fcea2b" stroke="#fcea2b" stroke-miterlimit="10"><circle cx="41" cy="37" r=".5"/><circle cx="40.2" cy="34.6" r=".5"/><circle cx="38.4" cy="32.9" r=".5"/><circle cx="36" cy="32.2" r=".5"/><circle cx="33.6" cy="32.9" r=".5"/><circle cx="31.8" cy="34.6" r=".5"/><circle cx="31" cy="37" r=".5"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}