package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HairDryerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHairDryerOne0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M11 12.138c0-.079.059-.146.137-.156L32.082 9.28C37.342 8.6 42 12.697 42 18c0 5.303-4.659 9.399-9.918 8.72l-20.945-2.702a.157.157 0 0 1-.137-.156V12.138Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M11 12L4 8v20l7-4"/><path d="m38 25l-6.694 15.898A5.07 5.07 0 0 1 26.634 44c-3.625 0-6.078-3.695-4.672-7.036L27 25"/><circle cx="35" cy="18" r="9" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHairDryerOne0)"/>`),
		g.Group(children),
	)
}