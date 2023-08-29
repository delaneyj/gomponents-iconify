package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlew(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm192-736q0-13-9-22.5t-22.5-9.5t-23 9.5T640 288v352q0 27-18.5 45.5T576 704h-48q-7 0-11.5-4.5T512 688V288q0-13-9.5-22.5T480 256t-22.5 9.5T448 288v400q0 7-4.5 11.5T432 704h-48q-26 0-45-19t-19-45V288q0-13-9-22.5t-22.5-9.5t-23 9.5T256 288v352q0 53 37.5 90.5T384 768h192q53 0 90.5-37.5T704 640V288z"/>`),
		g.Group(children),
	)
}