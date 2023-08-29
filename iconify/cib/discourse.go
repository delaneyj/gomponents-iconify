package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Discourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.135 0C24.885 0 32 7.313 32 15.995s-7.104 15.99-15.865 15.99L0 32.001V15.72C0 7.043 7.375.001 16.135.001zm.157 6.083a9.851 9.851 0 0 0-8.448 4.76a9.655 9.655 0 0 0-.198 9.625l-1.781 5.677l6.396-1.432a9.889 9.889 0 0 0 10.844-1.854a9.662 9.662 0 0 0 2.318-10.661a9.835 9.835 0 0 0-9.12-6.115z"/>`),
		g.Group(children),
	)
}