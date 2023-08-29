package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleyahoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512q0-104 40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199q0 139-68.5 257T769 955.5T512 1024zm320-688q0-7-4.5-11.5T816 320H592q-7 0-11.5 4.5T576 336v32q0 7 4.5 11.5T592 384h56L520 525L401 320h95q7 0 11.5-4.5T512 304v-32q0-7-4.5-11.5T496 256H272q-7 0-11.5 4.5T256 272v32q0 7 4.5 11.5T272 320h45v1l131 230v153h-48q-7 0-11.5 4.5T384 720v32q0 7 4.5 11.5T400 768h224q7 0 11.5-4.5T640 752v-32q0-7-4.5-11.5T624 704h-48V548l149-164h91q7 0 11.5-4.5T832 368v-32z"/>`),
		g.Group(children),
	)
}