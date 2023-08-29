package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Note(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M4 0C1.795 0 0 1.795 0 4v18c0 2.205 1.795 4 4 4h13c1.063 0 2.539-1.535 4.25-3.281c.24-.244.47-.473.719-.719c.246-.248.506-.51.75-.75C24.466 19.538 26 18.063 26 17V4c0-2.205-1.795-4-4-4H4zm0 2h18a2 2 0 0 1 2 2v13c0 .995-1.002 1-2 1h-3c-.551 0-1 .449-1 1v3.063c0 .887.002 1.753-.719 1.937H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2zm2.813 6A1.001 1.001 0 0 0 7 10h12a1 1 0 1 0 0-2H7a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0zm0 5A1.001 1.001 0 0 0 7 15h10a1 1 0 1 0 0-2H7a1 1 0 0 0-.094 0a1.001 1.001 0 0 0-.093 0z"/>`),
		g.Group(children),
	)
}