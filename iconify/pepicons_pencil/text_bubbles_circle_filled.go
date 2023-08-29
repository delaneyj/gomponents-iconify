package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextBubblesCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilTextBubblesCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M8.54 11.383a.833.833 0 1 0 0-1.666a.833.833 0 0 0 0 1.666Zm2.498 0a.833.833 0 1 0 0-1.666a.833.833 0 0 0 0 1.666Zm2.499 0a.833.833 0 1 0 0-1.666a.833.833 0 0 0 0 1.666Z"/><path d="M8.298 16.47c-.69 0-1.08-.87-1.28-1.58h-.15A2.173 2.173 0 0 1 4.7 12.72V8.47c.01-1.2.99-2.17 2.179-2.17h8.325c1.189 0 2.168.97 2.168 2.17v4.25c0 1.19-.97 2.17-2.168 2.17h-4.428c-.72.65-1.879 1.58-2.478 1.58ZM6.878 7.3c-.639 0-1.169.52-1.169 1.17v4.25c0 .64.51 1.17 1.16 1.17h.95l.08.4c.099.47.299 1 .439 1.17c.32-.13 1.17-.76 1.909-1.43l.14-.13h4.817c.64 0 1.169-.52 1.169-1.17V8.47c0-.64-.52-1.17-1.17-1.17H6.88Z"/><path d="M18.232 19.8c-.49 0-1.34-.66-2.019-1.27h-2.668a1.83 1.83 0 0 1-1.83-1.83c0-.28.22-.5.5-.5s.5.22.5.5c0 .46.37.83.83.83h3.058l.14.13c.53.48 1.139.95 1.419 1.09c.1-.17.23-.52.29-.83l.08-.4h.84c.459 0 .829-.37.829-.83v-3.4c0-.46-.37-.83-.83-.83c-.28 0-.5-.22-.5-.5s.22-.5.5-.5c1.01 0 1.829.82 1.829 1.83v3.4c0 1.01-.82 1.83-1.829 1.83h-.04c-.17.59-.5 1.27-1.1 1.27v.01Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilTextBubblesCircleFilled0)"/></g>`),
		g.Group(children),
	)
}