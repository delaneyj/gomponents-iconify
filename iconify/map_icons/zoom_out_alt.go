package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M35.66 29.539a18.044 18.044 0 0 0 2.632-9.385c0-10.029-8.115-18.15-18.146-18.154C10.124 2.003 2 10.125 2 20.152c0 10.018 8.125 18.139 18.152 18.139c3.44 0 6.645-.972 9.384-2.633L41.879 48L48 41.876L35.66 29.539zM20.15 31.38c-6.202-.015-11.216-5.027-11.227-11.216A11.245 11.245 0 0 1 20.15 8.935c6.199.016 11.215 5.028 11.228 11.229c-.013 6.182-5.031 11.201-11.228 11.216zM12 18h16v4H12z"/>`),
		g.Group(children),
	)
}