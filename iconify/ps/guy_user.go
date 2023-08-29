package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GuyUser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M3 448v64h320v-64q0-52-30-97.5T216 288v-23q29-15 46.5-43.5T280 160v-43q0-48-34.5-82.5T163 0T80 34.5T45 117v43q0 33 17.5 61.5T109 265v23q-46 17-76 62T3 448zm277 0v21H45v-21q0-36 19.5-67.5T116 333q14 30 47 30q31 0 47-30q31 16 50.5 48t19.5 67zm-45-341q-77-6-121-45q22-19 49-19t48 18.5t24 45.5zM88 117q0-9 4-21q61 48 145 53v11q0 31-22 53t-52 22q-31 0-53-22t-22-53v-43zm75 160h10v32q0 11-10 11q-11 0-11-11v-32h11z"/>`),
		g.Group(children),
	)
}