package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iitcmobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.077 9.009l-2.446 1.412v18.5l9.944 5.741"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m40.021 15.093l-8.042 4.644v18.157m-14.056 1.115l2.446-1.412v-18.5l-9.944-5.741"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.978 32.925l8.043-4.644V10.124"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.978 14.75v18.5L24 42.5l16.021-9.25v-18.5L24 5.5L7.978 14.75z"/>`),
		g.Group(children),
	)
}