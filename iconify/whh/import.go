package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Import(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M663 448h105q26 0 45 18.5t19 45t-19 45.5t-45 19H512q-27 0-45.5-19T448 512V256q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5v105L919 18q18-18 43.5-18t43.5 18t18 43.5t-18 43.5zM512 128q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512q0-27 18.5-45.5T960 448t45.5 18.5T1024 512q0 104-40.5 199t-109 163.5t-163.5 109t-199 40.5t-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0q26 0 45 18.5t19 45t-19 45.5t-45 19z"/>`),
		g.Group(children),
	)
}