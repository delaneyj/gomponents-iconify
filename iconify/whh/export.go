package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 384q-27 0-45.5-19T896 320V215L553 558q-18 18-43.5 18T466 558t-18-43.5t18-43.5l343-343H704q-27 0-45.5-19T640 63.5t18.5-45T704 0h256q27 0 45.5 18.5T1024 64v256q0 26-18.5 45T960 384zM512 128q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512q0-27 18.5-45.5T960 448t45.5 18.5T1024 512q0 104-40.5 199t-109 163.5t-163.5 109t-199 40.5t-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0q27 0 45.5 18.5t18.5 45t-19 45.5t-45 19z"/>`),
		g.Group(children),
	)
}