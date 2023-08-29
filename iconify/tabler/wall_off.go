package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h10a2 2 0 0 1 2 2v10m-.589 3.417c-.361.36-.86.583-1.411.583H6a2 2 0 0 1-2-2V6c0-.55.222-1.047.58-1.409M4 8h4m4 0h8m0 4h-4m-4 0H4m0 4h12M9 4v1m5 3v2m-6 2v4m3 0v4M3 3l18 18"/>`),
		g.Group(children),
	)
}