package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M19.5 17.796L14.019 13L19.5 8.204v9.592Zm-6.14-4.043a1 1 0 0 1 0-1.506l5.482-4.796c.646-.566 1.658-.106 1.658.753v9.592a1 1 0 0 1-1.659.753l-5.48-4.796Z"/><path d="M12.5 17.796L7.019 13L12.5 8.204v9.592Zm-6.14-4.043a1 1 0 0 1 0-1.506l5.481-4.796c.647-.566 1.659-.106 1.659.753v9.592a1 1 0 0 1-1.659.753l-5.48-4.796Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}