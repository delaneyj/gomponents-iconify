package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Male(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMale0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M41.952 15.048v-9h-9"/><path fill="#fff" d="M10.414 38c5.467 5.468 14.331 5.468 19.799 0a13.956 13.956 0 0 0 4.1-9.899a13.96 13.96 0 0 0-4.1-9.9c-5.468-5.467-14.332-5.467-19.8 0c-5.467 5.468-5.467 14.332 0 19.8Z"/><path stroke-linecap="round" d="m30 18l9.952-9.952"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMale0)"/>`),
		g.Group(children),
	)
}