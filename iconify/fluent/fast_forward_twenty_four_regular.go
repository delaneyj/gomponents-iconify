package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 5.503c0-1.28 1.498-1.971 2.472-1.142l7.41 6.306a1.75 1.75 0 0 1 0 2.665l-7.41 6.306c-.974.83-2.472.137-2.472-1.142v-3.987l-6.028 5.13C4.498 20.468 3 19.776 3 18.497V5.504c0-1.28 1.498-1.972 2.472-1.143l6.028 5.13V5.503Zm0 5.958l-7-5.957v12.993l7-5.957v-1.08Zm8.91.348L13 5.503v12.993l7.41-6.306a.25.25 0 0 0 0-.38Z"/>`),
		g.Group(children),
	)
}