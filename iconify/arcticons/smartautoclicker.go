package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartautoclicker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.707 33.206h-9.643a1.356 1.356 0 0 1-1.206-1.205V6.808a1.356 1.356 0 0 1 1.206-1.205h15.188a1.356 1.356 0 0 1 1.205 1.205v22.023"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.89 22.865l4.82 12.995l-4.42.584a2.008 2.008 0 0 0-1.165.653c-.327.359-1.037 1.127-.596 1.329l8.689 3.972l9.286-3.262a1.34 1.34 0 0 0 .613-1.454l-1.434-7.232a1.398 1.398 0 0 0-1.005-1.134l-7.747-.844l-.796.464l-2.678-7.238a2.115 2.115 0 0 0-2.408-1.082a1.999 1.999 0 0 0-1.159 2.249Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.093 26.877a5.183 5.183 0 1 1 7.52-2.887"/>`),
		g.Group(children),
	)
}