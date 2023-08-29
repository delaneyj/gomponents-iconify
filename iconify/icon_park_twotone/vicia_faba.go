package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViciaFaba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTViciaFaba0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M4 31c0-6.5 4-8.963 7-9.981C14 20 16 20 19 17s3-9 9-11s13.091 1 15 8c1.908 7-3.5 16-6 19s-7.501 8-16 9c-8.5 1-17-4.5-17-11Z"/><path stroke-linecap="round" d="M12 21.044c7 8.956 17 0 10-9.044"/><path stroke-linecap="round" d="M11 21.019C14 20 16 20 19 17s3-9 9-11"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTViciaFaba0)"/>`),
		g.Group(children),
	)
}