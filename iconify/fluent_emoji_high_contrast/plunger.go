package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plunger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.879 1.879A3 3 0 0 1 30.12 6.12L16.684 19.558a5.5 5.5 0 0 1-1.149 6.099l-.902.903a3.001 3.001 0 0 1-4.754 3.561l-8-8a3 3 0 0 1 3.561-4.754l.903-.902a5.501 5.501 0 0 1 6.099-1.15L25.879 1.88ZM15.586 15L17 16.414L28.707 4.707a1 1 0 0 0-1.414-1.414L15.586 15Zm-1.036 4.808l-2.376-2.366l.002-.002l-.08-.05a3.502 3.502 0 0 0-4.339.489l-1.524 1.524l6.365 6.363l1.523-1.523a3.502 3.502 0 0 0 .49-4.338l-.061-.097Zm.234-1.178L16 17.414L14.586 16l-1.219 1.219l1.417 1.411Zm-2.893 7.843L5.526 20.11l-.819-.817a1 1 0 0 0-1.414 1.414l8 8a1 1 0 0 0 1.414-1.414l-.818-.818l.002-.002Z"/>`),
		g.Group(children),
	)
}