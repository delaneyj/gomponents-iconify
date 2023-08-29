package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M23.964 3.98C12.918 3.99 3.97 12.951 3.98 23.997c.01 11.045 8.97 19.992 20.017 19.983c11.045-.009 19.992-8.97 19.983-20.016c-.009-11.046-8.97-19.993-20.016-19.984ZM19.99 33.047v-7h8v7h4v-18h-4v7h-8v-7h-4v18h4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}