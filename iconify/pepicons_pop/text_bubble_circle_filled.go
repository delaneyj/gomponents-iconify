package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBubbleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopTextBubbleCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g transform="translate(3 3)"><circle cx="7" cy="9" r="1" fill="#000"/><circle cx="10" cy="9" r="1" fill="#000"/><circle cx="13" cy="9" r="1" fill="#000"/><path fill-rule="evenodd" stroke="#000" stroke-linecap="round" stroke-width="2" d="M2 6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v6.102a2 2 0 0 1-2 2H8.45S6.364 16 5.705 16c-.66 0-1.056-1.898-1.056-1.898H4a2 2 0 0 1-2-2V6Z" clip-rule="evenodd"/><circle cx="7" cy="9" r="1" fill="#000"/><circle cx="10" cy="9" r="1" fill="#000"/><circle cx="13" cy="9" r="1" fill="#000"/><path fill="#000" fill-rule="evenodd" d="M1 6v6.102A3 3 0 0 0 3.864 15.1c.36 1.224.894 1.901 1.84 1.901c.757 0 1.684-.609 3.13-1.898H16a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H4a3 3 0 0 0-3 3Zm7.064 7.102l-.286.26c-.864.786-1.543 1.304-1.869 1.522a5.58 5.58 0 0 1-.283-.986l-.166-.796H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v6.102a1 1 0 0 1-1 1H8.064Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopTextBubbleCircleFilled0)"/></g>`),
		g.Group(children),
	)
}