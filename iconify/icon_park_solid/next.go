package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Next(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNext0"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M4 40.836c4.893-5.973 9.238-9.362 13.036-10.168c3.797-.805 7.412-.927 10.846-.365V41L44 23.545L27.882 7v10.167c-6.349.05-11.746 2.328-16.192 6.833C7.245 28.505 4.681 34.117 4 40.836Z" clip-rule="evenodd"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNext0)"/>`),
		g.Group(children),
	)
}