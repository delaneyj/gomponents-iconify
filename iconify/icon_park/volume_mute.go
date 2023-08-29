package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeMute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipVolumeMute0" width="13" height="13" x="30" y="18" maskUnits="userSpaceOnUse" style="mask-type:alpha"><rect width="13" height="13" x="30" y="18" fill="#fff"/></mask><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><g stroke-linecap="round" mask="url(#ipVolumeMute0)"><path d="M40.7348 20.2858L32.2495 28.7711"/><path d="M32.2496 20.2858L40.7349 28.7711"/></g><path fill="#2F88FF" d="M24 6V42C17 42 11.7985 32.8391 11.7985 32.8391H6C4.89543 32.8391 4 31.9437 4 30.8391V17.0108C4 15.9062 4.89543 15.0108 6 15.0108H11.7985C11.7985 15.0108 17 6 24 6Z"/></g>`),
		g.Group(children),
	)
}