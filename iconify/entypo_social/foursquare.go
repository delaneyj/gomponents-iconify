package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.92 1a.92.92 0 0 0-.92.92v16.516c0 .625.765.926 1.192.47l4.471-4.79A.357.357 0 0 1 9.927 14h3.237c.486 0 .905-.343 1.001-.82l2.111-10.514A1.392 1.392 0 0 0 14.911 1H4.92zm3.918 11.19L6 15.527V3.343C6 3.154 6.154 3 6.343 3h7.14a.54.54 0 0 1 .53.648L13.6 5.703a.369.369 0 0 1-.362.297h-3.71A.527.527 0 0 0 9 6.528v1.22c0 .139.113.252.253.252h3.294c.306 0 .536.28.476.581l-.614 3.058a.452.452 0 0 1-.442.361H9.25a.54.54 0 0 0-.412.19z"/>`),
		g.Group(children),
	)
}