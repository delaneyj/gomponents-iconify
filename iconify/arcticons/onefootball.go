package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onefootball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.941 7.926C16.297 40.171 16.53 39.862 16.53 39.862M35.726 8.06L24.621 40.075"/><path fill="none" stroke="currentColor" d="m16.53 39.861l8.092.213m3.319-32.148l7.785.135"/><ellipse cx="35.928" cy="33.403" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="6.572" ry="6.671"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.621 8.04c0 4.589-1.577 8.37-6.225 11.63c-4.399 3.085-9.83 3.057-12.895 1.605M24.621 8.04l-6.801-.046a6.32 6.32 0 0 1-3.36 5.871a9.565 9.565 0 0 1-6.616.875m0 0L5.5 21.277"/>`),
		g.Group(children),
	)
}