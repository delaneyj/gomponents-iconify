package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBat0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m21.619 8.986l.476 2.463h3.81l.476-2.463c.476 1.642 1.429 5.223 1.429 6.405c1.746-.164 5.143-1.38 4.762-4.927c0-.493-.382-1.676-1.905-2.464c4.285.986 12.952 5.716 13.333 16.754c-2.857-1.807-8.762-3.154-9.524 5.913c-1.746-2.957-5.809-7.293-8.095-.986C25.587 32.473 24 38.846 24 42c0-3.154-1.587-9.527-2.381-12.319c-2.286-6.307-6.35-1.97-8.095.986c-.762-9.067-6.667-7.72-9.524-5.913C4.381 13.716 13.048 8.986 17.333 8c-1.523.788-1.905 1.971-1.905 2.464c-.38 3.548 3.016 4.763 4.762 4.927c0-1.182.953-4.763 1.429-6.405Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBat0)"/>`),
		g.Group(children),
	)
}