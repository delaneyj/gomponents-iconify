package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDownloadFour0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#555" fill-rule="evenodd" stroke-linejoin="round" d="M24 29L12 17h8V6h8v11h8L24 29Z" clip-rule="evenodd"/><path d="M42 37H6m28 7H14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDownloadFour0)"/>`),
		g.Group(children),
	)
}