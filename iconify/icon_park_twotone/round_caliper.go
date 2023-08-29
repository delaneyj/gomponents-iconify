package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundCaliper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRoundCaliper0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill="#555" d="M24 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm20-4a20 20 0 0 1-9.284 16.887l-4.286-6.755A12 12 0 0 0 36 24h8ZM13.423 40.974A20 20 0 0 1 4 24.165l8-.066a12 12 0 0 0 5.654 10.086l-4.23 6.79Zm1.019-34.542a20 20 0 0 1 19.27.084l-3.885 6.994a12 12 0 0 0-11.562-.051l-3.823-7.027Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRoundCaliper0)"/>`),
		g.Group(children),
	)
}