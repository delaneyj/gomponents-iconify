package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkiffDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.92 3.43C10.508.217 1.712 15.697 3.163 26.996c2.18 16.973 22.342 21.966 25.905 3.945c-16.34-.875-18.032-24.781-.148-27.509V3.43Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.92 3.43c-32.874.453-16.924 47.309.148 27.509"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.08 44.57c18.412 3.213 27.208-12.267 25.757-23.565C42.657 4.031 22.495-.963 18.931 17.06c16.34.875 18.033 24.781.148 27.51h0v.001Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.08 44.57c32.874-.454 16.924-47.31-.148-27.51"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.274 26.605v-5.21h1.172c1.238 0 2.28 1.042 2.28 2.28v.65c0 1.238-1.042 2.28-2.28 2.28h-1.172Z"/>`),
		g.Group(children),
	)
}