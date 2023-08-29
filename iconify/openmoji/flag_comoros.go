package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagComoros(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 27h62v9H5z"/><path fill="#d22f27" d="M5 36h62v9H5z"/><path fill="#f1b31c" d="M5 17h62v10H5z"/><path fill="#5c9e31" d="M37 36L5 55V17l32 19z"/><g fill="#fff"><circle cx="19" cy="31" r="1.044"/><circle cx="19" cy="34.333" r="1.044"/><circle cx="19" cy="37.667" r="1.044"/><circle cx="19" cy="41" r="1.044"/></g><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M11.693 36.125a7.481 7.481 0 0 1 5.989-7.357a7.376 7.376 0 0 0-1.437-.143a7.5 7.5 0 0 0 0 15a7.384 7.384 0 0 0 1.437-.143a7.482 7.482 0 0 1-5.99-7.357Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}