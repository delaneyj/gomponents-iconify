package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autoit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 256C0 114.615 114.615 0 256 0s256 114.615 256 256s-114.615 256-256 256S0 397.385 0 256zM256 51C142.781 51 51 142.781 51 256s91.781 205 205 205s205-91.781 205-205S369.219 51 256 51z"/><path fill="currentColor" d="M221.301 162.6L101.181 332h68.6l87.92-124.32L310 283.84h-80.019l-35 48.16h217.84s-105.496-152.665-117.04-169.4s-21.195-19.772-35.618-20.856C234.915 139.846 221.3 162.6 221.3 162.6z"/>`),
		g.Group(children),
	)
}