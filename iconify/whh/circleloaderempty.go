package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleloaderempty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 896q122 0 225-60t163-163t60-225t-60-225T673 60T448 0T223 60T60 223T0 448t60 225t163 163t225 60zm0-832q104 0 192.5 51.5t140 140T832 448t-51.5 192.5t-140 140T448 832t-192.5-51.5t-140-140T64 448t51.5-192.5t140-140T448 64z"/>`),
		g.Group(children),
	)
}