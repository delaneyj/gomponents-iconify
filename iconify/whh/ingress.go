package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ingress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M480 1024L0 768V256L480 0l480 256v512zm416-736L512 81v175h-64V81L64 288v397l146-85l28 50l-161 93l403 217l404-217l-162-93l29-50l145 85V288zm-768 32h704L480 896zm352 174l174-110H307zm-189-51l157 270V544zm221 101v169l158-270z"/>`),
		g.Group(children),
	)
}