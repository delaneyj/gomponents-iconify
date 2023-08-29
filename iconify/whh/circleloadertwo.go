package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleloadertwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 512q0 91-35.5 174T765 829t-143 95.5T448 960t-174-35.5T131 829T35.5 686T0 512t35.5-174T131 195t143-95.5T448 64V0q104 0 199 40.5t163.5 109t109 163.5T960 512h-64zm-832 0q0 104 51.5 192.5t140 140T448 896t192.5-51.5t140-140T832 512h-64q0-87-43-160.5T608.5 235T448 192v-64q-104 0-192.5 51.5t-140 140T64 512z"/>`),
		g.Group(children),
	)
}