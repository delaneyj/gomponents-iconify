package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuslimPro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 18.358c0 5.18-9.847 6.716-9.847 14.441a8.384 8.384 0 0 0 2.044 5.224m7.803 0H6.12V42.5H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.094 42.5c0-2.473-.89-4.477-1.987-4.477S6.12 40.027 6.12 42.5m7.947 0c0-2.473-.89-4.477-1.987-4.477s-1.986 2.004-1.986 4.477"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.04 42.5c0-2.473-.89-4.477-1.987-4.477s-1.986 2.004-1.986 4.477m7.946 0c0-2.473-.89-4.477-1.986-4.477S18.04 40.027 18.04 42.5m7.947 0c0-2.473-.89-4.477-1.987-4.477s-1.987 2.004-1.987 4.477M11.358 21.152l1.024-3.687h-1.024v-.839c0-1.082-.878-4.048-1.96-4.048s-1.96 2.966-1.96 4.048v.84H6.412l1.024 3.686V25.3H6.413l1.024 3.688v9.035h3.921v-9.035l1.024-3.687h-1.024ZM24 18.358c0 5.18 9.847 6.716 9.847 14.441a8.384 8.384 0 0 1-2.044 5.224"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 38.023h17.88V42.5H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.907 42.5c0-2.473.889-4.477 1.986-4.477s1.987 2.004 1.987 4.477m-7.947 0c0-2.473.89-4.477 1.987-4.477s1.986 2.004 1.986 4.477m-7.946 0c0-2.473.89-4.477 1.987-4.477s1.986 2.004 1.986 4.477m-7.946 0c0-2.473.89-4.477 1.986-4.477s1.987 2.004 1.987 4.477m6.682-21.348l-1.025-3.687h1.025v-.839c0-1.082.877-4.048 1.96-4.048s1.96 2.966 1.96 4.048v.84h1.025l-1.024 3.686V25.3h1.024l-1.024 3.688v9.035h-3.921v-9.035L35.617 25.3h1.025ZM9.398 12.578v-1.872m29.204 1.872v-1.872m-12.397 5.301A5.562 5.562 0 0 1 23.643 5.5a6.434 6.434 0 1 0 6.187 9.137a5.53 5.53 0 0 1-3.625 1.37Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.692 6.725l.579 1.78h1.871l-1.514 1.1l.578 1.78l-1.514-1.1l-1.514 1.1l.579-1.78l-1.514-1.1h1.871l.578-1.78z"/>`),
		g.Group(children),
	)
}