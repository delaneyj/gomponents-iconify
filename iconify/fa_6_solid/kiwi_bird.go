package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KiwiBird(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M291.2 388.4c31.2-18.8 64.7-36.4 101.1-36.4H448c4.6 0 9.1-.2 13.6-.7l85.3 121.9c4 5.7 11.3 8.2 17.9 6.1S576 471 576 464V224c0-70.7-57.3-128-128-128h-55.7c-36.4 0-69.9-17.6-101.1-36.4C262.3 42.1 228.3 32 192 32C86 32 0 118 0 224c0 71.1 38.6 133.1 96 166.3V456c0 13.3 10.7 24 24 24s24-10.7 24-24v-46c15.3 3.9 31.4 6 48 6c5.4 0 10.7-.2 16-.7V456c0 13.3 10.7 24 24 24s24-10.7 24-24v-50.9c12.4-4.4 24.2-10 35.2-16.7zM448 200a24 24 0 1 1 0 48a24 24 0 1 1 0-48z"/>`),
		g.Group(children),
	)
}