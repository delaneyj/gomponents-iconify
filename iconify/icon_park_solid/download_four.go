package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDownloadFour0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke-linejoin="round" d="M24 29L12 17h8V6h8v11h8L24 29Z" clip-rule="evenodd"/><path d="M42 37H6m28 7H14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDownloadFour0)"/>`),
		g.Group(children),
	)
}