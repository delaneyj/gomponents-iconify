package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Friends(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 586q0 10-8.5 18t-26 13.5t-36 9.5t-47 6.5t-50 3.5t-54.5 2t-50.5 1H576V512h-1q-1-46-18-80q27-7 51-10v-31q-96-58-96-240q0-75 53.5-113T704 0t138.5 38T896 151q0 189-96 243v28q101 14 162 56t62 108zM416 778v28q101 14 162 56.5T640 970q0 10-8.5 18t-26 13.5t-36 9.5t-47 6.5t-50 4t-54.5 2t-50.5.5h-95l-50.5-.5l-54.5-2l-50-4l-47-6.5l-36-9.5l-26-13.5L0 970q1-65 62-107.5T224 806v-31q-96-58-96-240q0-75 53.5-113T320 384t138.5 38T512 535q0 189-96 243z"/>`),
		g.Group(children),
	)
}