package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forceps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-miterlimit="10" stroke-width="3.251" d="M23.21 63.42a4.609 4.609 0 0 0 3.514-2.904s22.26-40.97 25-51.13M19.33 61.22a4.956 4.956 0 0 1 .888-4.53s24.13-39.9 31.51-47.4"/><path d="m51.35 7.081l2.667 1.54a.975.975 0 0 1 .357 1.332l-5.254 9.101l-4.356-2.515l5.254-9.101a.975.975 0 0 1 1.332-.357z"/>`),
		g.Group(children),
	)
}