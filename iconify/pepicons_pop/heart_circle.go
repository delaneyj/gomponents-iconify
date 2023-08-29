package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.032 8.594C7.89 8.866 7 9.696 7 11.163c0 1.967 1.676 4.094 6 6.226c4.324-2.132 6-4.26 6-6.226c0-1.467-.889-2.297-2.032-2.57c-1.218-.29-2.523.103-3.167.965a1 1 0 0 1-1.602 0c-.644-.862-1.95-1.255-3.167-.964ZM13 7.543c-1.25-.98-2.965-1.245-4.432-.895C6.678 7.1 5 8.621 5 11.163c0 3.326 2.88 6.016 7.571 8.24a1 1 0 0 0 .857 0C18.12 17.18 21 14.49 21 11.164c0-2.542-1.678-4.064-3.568-4.515c-1.467-.35-3.183-.084-4.432.895Z"/><path d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z"/></g>`),
		g.Group(children),
	)
}