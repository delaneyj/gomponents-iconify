package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxTissue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M92.5 0H208c40 0 52 24 64 48s24 48 64 48h85.2c14.8 0 26.8 12 26.8 26.8c0 3.4-.7 6.8-1.9 10L409.6 224L384 288H128l-16-64L64.9 35.4c-.6-2.3-.9-4.6-.9-6.9C64 12.8 76.8 0 92.5 0zM79 224l16 64H80c-8.8 0-16 7.2-16 16s7.2 16 16 16h352c8.8 0 16-7.2 16-16s-7.2-16-16-16h-13.5l25.6-64H464c26.5 0 48 21.5 48 48v112H0V272c0-26.5 21.5-48 48-48h31zM0 416h512v48c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48v-48z"/>`),
		g.Group(children),
	)
}