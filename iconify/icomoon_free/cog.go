package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.59 9.535a3.053 3.053 0 0 1 1.127-4.164l-1.572-2.723a3.017 3.017 0 0 1-1.529.414A3.052 3.052 0 0 1 9.574 0H6.429a3.009 3.009 0 0 1-.406 1.535c-.839 1.454-2.706 1.948-4.17 1.106L.281 5.364a3 3 0 0 1 1.123 1.117a3.053 3.053 0 0 1-1.12 4.16l1.572 2.723c.448-.261.967-.41 1.522-.41A3.052 3.052 0 0 1 6.42 16h3.145a3.012 3.012 0 0 1 .406-1.519a3.053 3.053 0 0 1 4.163-1.11l1.572-2.723a3.008 3.008 0 0 1-1.116-1.113zM8 11.24a3.24 3.24 0 1 1 0-6.48a3.24 3.24 0 0 1 0 6.48z"/>`),
		g.Group(children),
	)
}