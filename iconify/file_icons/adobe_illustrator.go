package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeIllustrator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M172.049 108.008h53.334l76.693 256.941H255.75l-19.076-73.189h-74.746l-19.466 73.19h-40.877l70.465-244.873v-12.069zm54.113 145.99L199.689 145.38l-27.251 108.616h53.724zm155.333-80H339.45V364.95h42.045V174zM359.93 99.811c-13.621 0-24.663 10.915-24.663 24.38s11.042 24.381 24.663 24.381s24.664-10.915 24.664-24.38s-11.042-24.381-24.664-24.381zM512 512H0V0h512v512z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}