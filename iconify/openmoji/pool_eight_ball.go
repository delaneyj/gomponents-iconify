package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoolEightBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="35.958" r="29" fill="#3F3F3F"/><circle cx="36" cy="35.958" r="15" fill="#FFF"/><path fill-opacity=".6" d="M50.738 10.998c4.825 5.18 7.782 12.123 7.782 19.76c0 16.016-12.984 29-29 29a28.857 28.857 0 0 1-14.664-3.984C20.145 61.42 27.653 64.958 36 64.958c16.016 0 29-12.984 29-29c0-10.631-5.73-19.911-14.262-24.96z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="36" cy="35.958" r="29"/><circle cx="36" cy="32.595" r="3.363"/><circle cx="36" cy="40.31" r="4.352"/></g>`),
		g.Group(children),
	)
}