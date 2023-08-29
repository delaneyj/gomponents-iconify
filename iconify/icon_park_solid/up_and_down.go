package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpAndDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUpAndDown0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M12 36v8h32V12h-8v8h-8v8h-8v8h-8Z"/><path d="m18 13l9-9m-6 0h6v6M10 27H4v-6m9-3l-9 9"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUpAndDown0)"/>`),
		g.Group(children),
	)
}