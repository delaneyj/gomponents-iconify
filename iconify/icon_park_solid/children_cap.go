package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChildrenCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSChildrenCap0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 32c0-15 5-24 8-24s6 6 6 6h8s3-6 6-6s8 9 8 24"/><rect width="40" height="8" x="4" y="32" fill="#fff" rx="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSChildrenCap0)"/>`),
		g.Group(children),
	)
}