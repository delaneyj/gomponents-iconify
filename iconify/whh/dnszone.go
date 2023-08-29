package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dnszone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-896q-104 0-192.5 51.5t-140 140T128 512q0 84 34 158t94 128V320q0-27 19-45.5t45.5-18.5t45 18.5T384 320v554q63 22 128 22q104 0 192.5-51.5t140-140T896 512t-51.5-192.5t-140-140T512 128zm192 448L512 704q-26 0-45-19t-19-45V384q0-27 19-45.5t45-18.5l192 128q1 0 12.5 6t15 8.5T744 471t13 11.5t7.5 13.5t3.5 16q0 6-1.5 12t-4 10.5t-7.5 9.5l-8 8l-10.5 7l-10 5.5l-11.5 6z"/>`),
		g.Group(children),
	)
}