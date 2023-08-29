package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sbt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M162.818 317.784h341.8c2.95-12.01 5.03-24.434 6.221-37.206H224.26c-39.143 0-39.143-59.393 0-59.393h285.435a252.802 252.802 0 0 0-7.95-37.547H285.699c-40.236 0-40.236-59.393 0-59.393h189.892C407.269 11.671 255.233-39.065 127.713 34.452c-170.284 98.17-170.284 344.929 0 443.099c131.445 75.78 288.918 19.527 353.925-100.375h-318.82c-39.333 0-39.333-59.392 0-59.392z"/>`),
		g.Group(children),
	)
}