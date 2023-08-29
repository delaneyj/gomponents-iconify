package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feelgood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFeelgood0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M23.142 15.858L11.828 27.172a4 4 0 0 0 0 5.656l9.9 9.9a4 4 0 0 0 5.657 0l9.9-9.9a4 4 0 0 0 0-5.656L25.97 15.858a2 2 0 0 0-2.829 0Z"/><path d="M25 13.84s4.554-5.11 6.373-7.266c1.818-2.156 3.768-1.878 5.392-.476c1.625 1.402 1.668 3.616 0 5.677C35.097 13.835 30.556 19 30.556 19M25 13.839s-4.554-5.11-6.373-7.266c-1.818-2.156-3.767-1.878-5.392-.476c-1.625 1.403-1.668 3.617 0 5.677C14.903 13.835 19.444 19 19.444 19"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFeelgood0)"/>`),
		g.Group(children),
	)
}