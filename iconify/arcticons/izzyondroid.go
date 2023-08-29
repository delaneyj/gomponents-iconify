package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Izzyondroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.145 6.048L20.973 17.54l13.66.022M17.196 23.38L5.607 41.723"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.634 17.562a12.6 12.6 0 0 1-5.01 24.161H5.608m7.205-11.514a12.6 12.6 0 0 1 5.01-24.161H41.84L31.578 22.292m-4.855 7.684l-7.421 11.747"/><path fill="currentColor" d="M21.57 24.315a.75.75 0 1 1-.75.75a.75.75 0 0 1 .75-.75Zm5.247 0a.752.752 0 1 1-.004 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.309 20.519l-1.889 1.829m-8.64-1.83l1.893 1.825m9.544 5.174a6.173 6.173 0 0 0-12.345 0Z"/>`),
		g.Group(children),
	)
}