package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsAddInTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.5 3A2.5 2.5 0 0 1 13 5.5V11h5.5a2.5 2.5 0 0 1 2.5 2.5v5a2.5 2.5 0 0 1-2.5 2.5h-13A2.5 2.5 0 0 1 3 18.5v-13A2.5 2.5 0 0 1 5.5 3h5Zm.5 10H5v5.5a.5.5 0 0 0 .5.5H11v-6Zm7.5 0H13v6h5.5a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5Zm-8-8h-5a.5.5 0 0 0-.5.5V11h6V5.5a.5.5 0 0 0-.5-.5Zm7.383-2.993L18 2a1 1 0 0 1 .993.883L19 3v2h2a1 1 0 0 1 .993.883L22 6a1 1 0 0 1-.883.993L21 7h-2v2a1 1 0 0 1-.883.993L18 10a1 1 0 0 1-.993-.883L17 9V7h-2a1 1 0 0 1-.993-.883L14 6a1 1 0 0 1 .883-.993L15 5h2V3a1 1 0 0 1 .883-.993L18 2l-.117.007Z"/>`),
		g.Group(children),
	)
}