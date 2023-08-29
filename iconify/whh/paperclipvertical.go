package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclipvertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M351.5 1024q-95.5 0-176.5-46.5T47 851T0 676V256Q0 150 75 75T256 0t181 75t75 181v352q0 66-47 113t-113 47t-113-47t-47-113V448q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v160q0 13 9.5 22.5T352 640t22.5-9.5T384 608V256q0-128-128-128T128 256v410q0 95 65.5 162.5T352 896t158.5-67.5T576 666V384q0-27 18.5-45.5t45-18.5t45.5 19t19 45v292q0 95-47.5 175T528 977.5T351.5 1024z"/>`),
		g.Group(children),
	)
}