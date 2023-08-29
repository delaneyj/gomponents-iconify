package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Form(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSForm0"><g fill="none"><rect width="40" height="32" x="4" y="8" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="2"/><path fill="#fff" fill-rule="evenodd" d="M4 29h40H4Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 29h40"/><path fill="#fff" fill-rule="evenodd" d="M4 19h40H4Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 19h40"/><path fill="#fff" fill-rule="evenodd" d="M17 40V19v21Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 40V19"/><path fill="#fff" fill-rule="evenodd" d="M4 38V17v21Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 38V17"/><path fill="#fff" fill-rule="evenodd" d="M44 38V17v21Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 38V17"/><path fill="#fff" fill-rule="evenodd" d="M31 40V19v21Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 40V19"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 40h30"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSForm0)"/>`),
		g.Group(children),
	)
}