package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mapmarker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m657 680l-272 345l-272-345q-73-76-99.5-178.5t0-205t100-179T286 14t198 0t172.5 103.5t100 179t0 205T657 680zM385.5 130Q279 130 204 205t-75 181t75 180.5T385.5 641T567 566.5T642 386t-75-181t-181.5-75z"/>`),
		g.Group(children),
	)
}