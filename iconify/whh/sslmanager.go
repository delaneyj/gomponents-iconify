package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sslmanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384 1024q-104 0-192.5-51.5t-140-140T0 640q0-116 64-212V320q0-87 43-160.5T223.5 43T384 0t160.5 43T661 159.5T704 320v108q64 96 64 212q0 104-51.5 192.5t-140 140T384 1024zm0-896q-76 0-131 52t-60 127q89-51 191-51t191 51q-5-75-60-127t-131-52zm0 256q-106 0-181 75t-75 181t75 181t181 75t181-75t75-181t-75-181t-181-75zm0 384q-53 0-90.5-37.5T256 640q0-64 52-102l76-153l76 153q52 38 52 102q0 53-37.5 90.5T384 768z"/>`),
		g.Group(children),
	)
}