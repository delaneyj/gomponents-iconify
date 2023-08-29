package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M240.015 89.722L0 734.546l96.021 375.732H1103.98L1200 734.546L959.985 89.722h-719.97zM93.75 793.945h1012.573l-65.625 256.934H159.375L93.75 793.945zm834.229 80.493c-26.508 0-47.974 21.466-47.974 47.974c0 26.509 21.466 48.047 47.974 48.047s48.047-21.538 48.047-48.047c-.001-26.508-21.54-47.974-48.047-47.974z"/>`),
		g.Group(children),
	)
}