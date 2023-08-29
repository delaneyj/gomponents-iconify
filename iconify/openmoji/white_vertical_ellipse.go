package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteVerticalEllipse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M56 36c0 16.016-8.954 29-20 29S16 52.016 16 36S24.954 7 36 7s20 12.984 20 29z"/><path fill="none" stroke="#000" stroke-width="2" d="M56 35.957c.035 16.016-8.891 29.02-19.937 29.044C25.017 65.025 16.034 52.06 16 36.043c-.035-16.016 8.891-29.02 19.937-29.044C46.983 6.975 55.966 19.94 56 35.957z"/>`),
		g.Group(children),
	)
}