package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMailDownload0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 18H4v24h40V18h-6"/><path fill="#fff" d="M38 6H10v16.5L24 33l14-10.5V6Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 22.5L24 33l14-10.5m-28 0V6h28v16.5m-28 0L4 18m34 4.5l6-4.5"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m19 19l5 5m0 0l5-5m-5 5V13"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMailDownload0)"/>`),
		g.Group(children),
	)
}