package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleArrowLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2a10 10 0 0 1 .324 19.995L12 22l-.324-.005A10 10 0 0 1 12 2zm.707 5.293a1 1 0 0 0-1.414 0l-4 4a1.048 1.048 0 0 0-.083.094l-.064.092l-.052.098l-.044.11l-.03.112l-.017.126L7 12l.004.09l.007.058l.025.118l.035.105l.054.113l.043.07l.071.095l.054.058l4 4l.094.083a1 1 0 0 0 1.32-1.497L10.415 13H16l.117-.007A1 1 0 0 0 16 11h-5.586l2.293-2.293l.083-.094a1 1 0 0 0-.083-1.32z"/></g>`),
		g.Group(children),
	)
}