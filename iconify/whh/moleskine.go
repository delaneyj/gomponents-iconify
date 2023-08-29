package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moleskine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768.338 1024h-64V704h128V320h-128V0h64q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-704-704v384h576v320h-512q-53 0-90.5-37.5T.338 896V128q0-53 37.5-90.5t90.5-37.5h512v320h-576z"/>`),
		g.Group(children),
	)
}