package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sapling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSapling0"><g fill="none"><path fill="#555" d="M27 8c-4.686 4.686-3.965 13.035-3.965 13.035S31.314 21.686 36 17s4.006-12.965 4.006-12.965S31.686 3.314 27 8Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M23.035 21.035L31 13.07m-7.965 7.965S22.314 12.686 27 8s13.006-3.965 13.006-3.965S40.686 12.314 36 17s-12.965 4.035-12.965 4.035Z"/><path fill="#555" d="M20 21c3.85 3.843 2.971 10.971 2.971 10.971s-7.122.872-10.97-2.971c-3.85-3.843-2.969-10.971-2.969-10.971S16.151 17.157 20 21Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M22.971 31.971L16.001 25m6.97 6.971s.878-7.128-2.97-10.971c-3.85-3.843-10.969-2.971-10.969-2.971S8.151 25.157 12 29c3.85 3.843 10.971 2.971 10.971 2.971ZM23 4v40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSapling0)"/>`),
		g.Group(children),
	)
}