package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneModeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l23 23m-8.999-8.999c3.74.043 8.999.999 8.999.999v-2.25S18 10.5 14.25 9.5V4.211c0-.79-.23-1.565-.707-2.194A15.368 15.368 0 0 0 12.25.5h-.5a15.17 15.17 0 0 0-1.293 1.517c-.478.629-.707 1.404-.707 2.194V9.5l-.196.054m4.33 7.83A46.418 46.418 0 0 0 13.75 21s1.25.75 2.25 1.75v.5S13.75 23 12 23s-4 .25-4 .25v-.5c1-1 2.25-1.75 2.25-1.75c0-4.22-.5-6.5-.5-6.5c-3.75 0-9.25 1-9.25 1v-2.25s3.232-1.616 6.43-2.82"/>`),
		g.Group(children),
	)
}