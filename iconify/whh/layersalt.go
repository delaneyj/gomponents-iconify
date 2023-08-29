package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layersalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M985 378L607 615q-40 24-95 24t-95-24L39 378Q0 354 0 319.5T39 260L417 24q40-24 95-24t95 24l378 236q39 25 39 59.5T985 378zM39 644l111-69l267 168q40 24 95 24t95-24l267-168l111 69q39 25 39 59.5T985 762L607 999q-40 24-95 24t-95-24L39 762Q0 738 0 703.5T39 644z"/>`),
		g.Group(children),
	)
}