package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KakaoSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1792 1652V140q0-58-41-99t-99-41H140Q82 0 41 41T0 140v1512q0 58 41 99t99 41h1512q58 0 99-41t41-99zM896 252q198 0 365.5 77t265 209t97.5 288t-97.5 288t-265 209t-365.5 77q-66 0-129-9q-68 49-180.5 125T459 1594q-9 4-17-1q-4-2-5.5-6.5t-1.5-7.5v-3q1-5 74-264q-156-77-248.5-206T168 826q0-156 97.5-288t265-209T896 252z"/>`),
		g.Group(children),
	)
}