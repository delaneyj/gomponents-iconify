package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUsb0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M12 22a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm24 6a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="m19 9l5-5l5 5m-4 30L12 28.263V22m24 6v4.79L24 41m0-37v39m-3 1h6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUsb0)"/>`),
		g.Group(children),
	)
}