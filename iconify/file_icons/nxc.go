package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nxc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M342.352 55.717H0V459.32h342.352V55.717zM57.657 95.005h223.68V242.95H57.658V95.005zm46.152 272.558L42.27 330.958l61.54-38.12v74.725zm99.636 66.074h-71.04v-42.792h71.04v42.792zm0-65.95h-71.04v-73.643h71.04v73.643zm31.74-.124v-74.726l61.539 38.121l-61.539 36.605zM251.45 0h73.428v33.591H251.45V0zM19.51 0h212.745v33.591H19.51V0zm305.368 480.446V512H17.474v-31.554h307.404z"/>`),
		g.Group(children),
	)
}